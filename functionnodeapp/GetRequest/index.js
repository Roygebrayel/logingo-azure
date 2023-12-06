const sql = require("mssql");

const config = {
  user: "roy",
  password: "ROge313313313313",
  server: "sqlservergo.database.windows.net",
  database: "roysqldb",
  options: {
    encrypt: true, // Use this if you're connecting to Azure SQL Database
  },
};

module.exports = async function (context, req) {
  try {
    // Create a connection pool
    const pool = await sql.connect(config);

    // Execute the query
    const result = await pool
      .request()
      .query("SELECT username, password FROM users");

    // Map the result to an array of users
    const users = result.recordset.map((row) => ({
      username: row.username,
      password: row.password,
    }));

    // Respond with the list of users
    context.res = {
      status: 200,
      body: JSON.stringify(users),
    };
  } catch (error) {
    // Handle errors
    context.res = {
      status: 500,
      body: JSON.stringify({ error: error.message }),
    };
  } finally {
    // Close the connection pool
    sql.close();
  }
};
