module.exports = async function (context, myTimer) {
  if (myTimer.isPastDue) {
    context.log("JavaScript is running late!");
  }

  context.log("Hello, I'm the trigger!");

  // Display the HTML message
  context.res = {
    headers: {
      "Content-Type": "text/html",
    },
    body: "<html><body><h1>Hello, I'm the trigger!</h1></body></html>",
  };

  // Set a timeout to clear the message after 5 seconds
  setTimeout(() => {
    context.res = {
      status: 204,
      body: "",
    };
    context.log("Message cleared.");
    context.done();
  }, 5000);
};
