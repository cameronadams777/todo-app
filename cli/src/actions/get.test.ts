import axios, { AxiosResponse } from "axios";
import { getTodo } from "./get";

jest.mock("axios", () => ({
  get: jest.fn()
}));

describe("getTodos", () => {
  beforeEach(() => jest.resetAllMocks());

  it("should call all todos endpoint if --all arg is true", async () => {
    (axios.get as jest.Mock).mockImplementation(jest.fn);
    await getTodo({ all: true });
    expect(axios.get).toHaveBeenCalled();
    expect(axios.get).toHaveBeenCalledWith("http://localhost:5000/todos");
  });

  it("should call todo by id endpoint with id that was passed via --id arg", async () => {
    (axios.get as jest.Mock).mockImplementation(jest.fn);
    await getTodo({ all: false, id: 8 });
    expect(axios.get).toHaveBeenCalled();
    expect(axios.get).toHaveBeenCalledWith("http://localhost:5000/todos/8");
  });

  it("should call todo by id endpoint if --all arg is true but id is still passed", async () => {
    (axios.get as jest.Mock).mockImplementation(jest.fn);
    await getTodo({ all: true, id: 8 });
    expect(axios.get).toHaveBeenCalled();
    expect(axios.get).toHaveBeenCalledWith("http://localhost:5000/todos/8");
  });

  it("should throw an error if --all arg is false but --id arg is not provided", async () => {
    await expect(getTodo({ all: false })).rejects.toThrow(
      "Todo id must be specified if --all params is false."
    );
  });

  it("should output the response of the get request when hitting the get all endpoint", async () => {
    const response: Partial<AxiosResponse> = {
      data: [
        {
          ID: 8,
          CreatedAt: "2022-06-20T21:34:11.664963-05:00",
          UpdatedAt: "2022-06-20T21:34:11.664963-05:00",
          DeletedAt: null,
          Title: "walk the dog",
          Description: "",
          CompletedAt: null,
          UserID: 1
        }
      ]
    };
    const loggerSpy = jest.spyOn(console, "log").mockImplementation(jest.fn);
    (axios.get as jest.Mock).mockResolvedValue(response);
    await getTodo({ all: true });
    expect(axios.get).toHaveBeenCalledWith("http://localhost:5000/todos");
    expect(loggerSpy).toHaveBeenCalledWith(response.data);
  });

  it("should output the response of the get request when hitting the get todo by id endpoint", async () => {
    const response: Partial<AxiosResponse> = {
      data: {
        ID: 8,
        CreatedAt: "2022-06-20T21:34:11.664963-05:00",
        UpdatedAt: "2022-06-20T21:34:11.664963-05:00",
        DeletedAt: null,
        Title: "walk the dog",
        Description: "",
        CompletedAt: null,
        UserID: 1
      }
    };
    const loggerSpy = jest.spyOn(console, "log").mockImplementation(jest.fn);
    (axios.get as jest.Mock).mockResolvedValue(response);
    await getTodo({ all: true, id: 8 });
    expect(axios.get).toHaveBeenCalledWith("http://localhost:5000/todos/8");
    expect(loggerSpy).toHaveBeenCalledWith(response.data);
  });
});
