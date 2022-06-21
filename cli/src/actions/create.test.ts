import axios from "axios";
import colors from "colors";
import { createTodo } from "./create";

jest.mock("axios", () => ({
  post: jest.fn()
}));

jest.mock("colors", () => ({
  red: jest.fn(),
  green: jest.fn()
}));

describe("createTodos", () => {
  beforeEach(() => jest.clearAllMocks());

  it("should throw error if body argument is not provided", async () => {
    await expect(createTodo({})).rejects.toThrow(
      "Body argument must be passed"
    );
  });

  it("should throw error if body argument does not contain a title", async () => {
    await expect(createTodo({ body: JSON.stringify({}) })).rejects.toThrow(
      "Todo title must be passed in the body."
    );
  });

  it("should throw error if body argument does not contain a user id", async () => {
    await expect(
      createTodo({ body: JSON.stringify({ title: "Walk the dog" }) })
    ).rejects.toThrow("User ID must be passed in the body.");
  });

  it("should make post request with correct parameters given that the body param is valid", async () => {
    const expectedBody = JSON.stringify({ title: "Walk the dog", userId: 8 });
    (axios.post as jest.Mock).mockImplementation(jest.fn);
    await createTodo({ body: expectedBody });
    expect(axios.post).toHaveBeenCalled();
    expect(axios.post).toHaveBeenCalledWith(
      "http://localhost:5000/todos",
      expectedBody
    );
  });

  it("should call console.log with the response from the api if the request is successful", async () => {
    const loggerSpy = jest.spyOn(console, "log").mockImplementation(jest.fn);
    (axios.post as jest.Mock).mockResolvedValue({
      status: 201,
      data: { message: "SUCCESS" }
    });
    (colors.green as unknown as jest.Mock).mockImplementation(
      (arg: any) => arg
    );
    await createTodo({
      body: JSON.stringify({ title: "Walk the dog", userId: 8 })
    });
    expect(loggerSpy).toHaveBeenCalledWith({
      message: "SUCCESS"
    });
  });
});
