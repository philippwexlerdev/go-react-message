import { User } from "../types/User";

export const getUserInfo = () =>
  JSON.parse(JSON.parse(localStorage.getItem("token") || "{}") || "{}") as User;
