import axios from "axios";
import colors from "colors";
import { completeTodo } from "./complete";

jest.mock("axios", () => ({
  post: jest.fn()
}));

jest.mock("colors", () => ({
  red: jest.fn(),
  green: jest.fn()
}));

describe("completeTodos", () => {
  it("should throw error if id argument is not provided", async () => {
    await expect(completeTodo({})).rejects.toThrow(
      "ID argument must be passed"
    );
  });

  it("should make post request with id parameter in the body", async () => {
    const expectedBody = { id: 8 };
    (axios.post as jest.Mock).mockImplementation(jest.fn);
    await completeTodo({ id: 8 });
    expect(axios.post).toHaveBeenCalled();
    expect(axios.post).toHaveBeenCalledWith(
      "http://localhost:5000/todos/complete",
      expectedBody
    );
  });

  it("should call console.log with the response from the api if the request is successful", async () => {
    const loggerSpy = jest.spyOn(console, "log").mockImplementation(jest.fn);
    (axios.post as jest.Mock).mockResolvedValue({
      status: 200,
      data: { message: "SUCCESS" }
    });
    (colors.green as unknown as jest.Mock).mockImplementation(
      (arg: any) => arg
    );
    await completeTodo({ id: 8 });
    expect(loggerSpy).toHaveBeenCalledWith({
      message: "SUCCESS"
    });
  });
});
