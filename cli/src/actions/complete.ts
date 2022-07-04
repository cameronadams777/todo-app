import axios from "axios";
import colors from "colors";
import ora from "ora-classic";

export const completeTodo = async (args: Record<string, any>) => {
  if (!args.id) throw new Error("ID argument must be passed");

  try {
    const spinner = ora("Completing todo...").start();

    const response = await axios.post("http://localhost:5000/todos/complete", {
      id: args.id
    });

    if (response.status === 200) spinner.prefixText = colors.green("COMPLETED");
    else spinner.prefixText = colors.red("FAILED");

    spinner.stopAndPersist();
    console.log(colors.green(response.data));
  } catch (error) {
    console.error(
      colors.red(
        `The following error occurred while completing your todo item: ${error}`
      )
    );
  }
};
