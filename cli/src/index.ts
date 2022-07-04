#!/usr/bin/env node
import { program } from "commander";
import pkg from "../package.json";
import { completeTodo } from "./actions/complete";
import { createTodo } from "./actions/create";
import { getTodo } from "./actions/get";

program.name("todos").description(pkg.description).version(pkg.version);

program
  .command("create")
  .option("-t, --title <title>", "the title of the todo item to be created")
  .option("-u, --userId <userId>", "the id of the user creating the todo item")
  .action(createTodo);

program
  .command("get")
  .option("-a, --all", "whether to get all todo items or not", true)
  .option("-i, --id", "the specific id of the todo item we want to get")
  .action(getTodo);

program
  .command("complete")
  .option("-i, --ids", "list of comma separated ids to complete")
  .action(completeTodo);

program.parse();
