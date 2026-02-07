import type { Tool } from "./types/tool";
import { mathTools } from "./tools/math";
import { weatherTools } from "./tools/weather";

export const tools: Tool[] = [...mathTools, ...weatherTools];

export function findTool(name: string) {
  return tools.find((t) => t.name === name);
}
