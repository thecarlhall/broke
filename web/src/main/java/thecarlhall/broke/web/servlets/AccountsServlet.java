/**
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package thecarlhall.broke.web.servlets;

import static com.google.common.base.Preconditions.checkNotNull;
import static javax.servlet.http.HttpServletResponse.SC_BAD_REQUEST;
import static javax.servlet.http.HttpServletResponse.SC_INTERNAL_SERVER_ERROR;
import static javax.servlet.http.HttpServletResponse.SC_NOT_FOUND;
import static javax.servlet.http.HttpServletResponse.SC_OK;

import java.io.IOException;
import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.SQLException;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import com.fasterxml.jackson.core.JsonGenerator;

public class AccountsServlet extends BaseServlet {
    private static final long serialVersionUID = 1L;

    private static final Logger LOGGER = LoggerFactory
            .getLogger(AccountsServlet.class);

    /**
     * @see javax.servlet.http.HttpServlet#doGet(javax.servlet.http.HttpServletRequest,
     *      javax.servlet.http.HttpServletResponse)
     */
    @Override
    protected void doGet(HttpServletRequest req, HttpServletResponse resp)
            throws ServletException, IOException {
        try {
            checkNotNull(req.getParameter("account_id"), "account_id is required");
        } catch (Exception e) {
            resp.sendError(SC_BAD_REQUEST, e.getMessage());
        }

        Connection conn = getReadConnection();
        try {
            int accountId = Integer.parseInt(req.getParameter("account_id"));
            PreparedStatement ps = conn.prepareStatement("select accounts.*, frequency.key as freq_key from accounts inner join frequency on accounts.frequency_id = frequency.id where accounts.id = ? limit 1");
            ps.setInt(1, accountId);
            ResultSet rs = ps.executeQuery();

            // output the first row if data was found
            if (rs.next()) {
                JsonGenerator writer = createJsonWriter(resp);

                writer.writeStartObject();
                writer.writeStringField("name", rs.getString("name"));
                writer.writeNumberField("dueDate", rs.getInt("due_date"));
                writer.writeStringField("frequency", rs.getString("freq_key"));
                writer.writeEndObject();
                writer.close();
                resp.setStatus(SC_OK);
            } else {
                resp.sendError(SC_NOT_FOUND,
                        String.format("Account ID [%s] not found", accountId));
            }
        } catch (SQLException e) {
            LOGGER.error(e.getMessage(), e);
            resp.sendError(SC_INTERNAL_SERVER_ERROR, e.getMessage());
        } finally {
            closeConnection(conn);
        }
    }
}