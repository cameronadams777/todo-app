import axios from "axios";
import colors from "colors";
import ora from "ora-classic";

export const createTodo = async (args: Record<string, any>) => {
  if (!args.body) throw new Error("Body argument must be passed.");
  if (!JSON.parse(args.body).title)
    throw new Error("Todo title must be passed in the body.");
  if (!JSON.parse(args.body).userId)
    throw new Error("User ID must be passed in the body.");
  const spinner = ora("Creating todo...").start();
  const { status } = await axios.post("http://localhost:5000/todo", args.body);
  if (status === 201) spinner.prefixText = colors.green("CREATED");
  else spinner.prefixText = colors.red("FAILED");
  spinner.stopAndPersist();
};
