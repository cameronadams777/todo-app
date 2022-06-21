import axios from "axios";
import colors from "colors";

const getAllTodoItems = async (): Promise<void> => {
  try {
    const response = await axios.get("http://localhost:5000/todos");
    console.log(response.data);
  } catch (error) {
    console.error(
      colors.red(
        `The following error occurred while retrieving todo items: ${error}`
      )
    );
  }
};

const getTodoItemById = async (id: string): Promise<void> => {
  try {
    const response = await axios.get(`http://localhost:5000/todos/${id}`);
    console.log(response.data);
  } catch (error) {
    console.error(
      colors.red(
        `The following error occurred while retrieving todo items: ${error}`
      )
    );
  }
};

export const getTodo = async (args: Record<string, any>): Promise<void> => {
  // Assume that if id argument is not null then we don't want all arguments
  const shouldRetrieveAllTodoItems = args.all && args.id == null;

  if (shouldRetrieveAllTodoItems) {
    await getAllTodoItems();
    return;
  }

  if (!args.id)
    throw new Error("Todo id must be specified if --all params is false.");

  await getTodoItemById(args.id);
};
