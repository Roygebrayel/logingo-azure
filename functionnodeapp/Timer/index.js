module.exports = async function (context, Timer) {
  // Validate timer object
  if (Timer.isPastDue) {
    context.log("Timer function ran late!");
  }

  // Log a message to the console
  context.log("Timer trigger function ran!", new Date().toISOString());

  // Replace the log statement with your custom logic

  // Complete the function execution
  context.done();
};
