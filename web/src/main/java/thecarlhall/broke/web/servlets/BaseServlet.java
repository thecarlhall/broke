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

import java.io.IOException;
import java.sql.Connection;
import java.sql.SQLException;

import javax.naming.InitialContext;
import javax.naming.NamingException;
import javax.servlet.ServletResponse;
import javax.servlet.http.HttpServlet;
import javax.sql.DataSource;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import com.fasterxml.jackson.core.JsonFactory;
import com.fasterxml.jackson.core.JsonGenerator;

public abstract class BaseServlet extends HttpServlet {
    private static final JsonFactory JSON_FACTORY = new JsonFactory();
    private static final String DATA_SOURCE_NAME = "java:comp/env/jdbc/BrokeDS";
    private static final Logger LOGGER = LoggerFactory.getLogger(BaseServlet.class);

    /**
     * Get a connection to the database for read purposes.
     *
     * @return A {@link Connection} to the database for reading.
     */
    public Connection getReadConnection() {
        try {
            InitialContext ic = new InitialContext();
            DataSource ds = (DataSource) ic.lookup(DATA_SOURCE_NAME);
            Connection conn = ds.getConnection();
            return conn;
        } catch (NamingException | SQLException e) {
            LOGGER.error(e.getMessage(), e);
            throw new RuntimeException(e.getMessage(), e);
        }
    }
    
    /**
     * Get a connection to the database for write purposes.
     *
     * @return A {@link Connection} to the database for writing.
     */
    public Connection getWriteConnection() {
        return getReadConnection();
    }

    /**
     * Close connection to the database. This saves us from having to catch
     * exceptions when we don't care if closing the connection errors out.
     *
     * @param conn
     *            The {@link Connection} to close.
     */
    public void closeConnection(Connection conn) {
        try {
            conn.close();
        } catch (SQLException e) {
            // meh
        }
    }

    /**
     * Create JSON writer that uses the HTTP servlet response.
     *
     * @param resp
     * @return
     * @throws IOException
     */
    public JsonGenerator createJsonWriter(ServletResponse resp)
            throws IOException {
        return JSON_FACTORY.createGenerator(resp.getOutputStream());
    }
}