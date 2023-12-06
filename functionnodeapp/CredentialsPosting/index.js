const isValidLogin = (password) => {
  // Ensure the password is at least 8 characters long
  return password.length >= 8;
};

module.exports = async function (context, req) {
  context.log("CredentialsPosting function processed a request.");

  if (req.method === "POST") {
    const { username, password } = req.body;

    // Validate login credentials
    if (!isValidLogin(password)) {
      // Render a failed login response
      context.res = {
        status: 400,
        body: "Invalid credentials",
      };
    } else {
      // Handle successful login
      context.res = {
        status: 200,
        body: `Welcome, ${username}!`,
      };
    }
  } else {
    // Handle other HTTP methods
    context.res = {
      status: 405,
      body: "Method Not Allowed",
    };
  }
};
