import axios from "../config/axios";
import { getUserInfo } from "../util/getUserInfo";

interface LoginPayload {
  username: string;
  password: string;
}

export const loginUser = async ({ username, password }: LoginPayload) => {
  try {
    const data = await axios.post("/user/login", { username, password });
    if (data.status !== 200) {
      throw data;
    }
    return data.data;
  } catch (err: any) {
    if (err?.response?.data?.status === 200) {
      return err?.response?.data;
    }
    throw err;
  }
};

interface SignupPayload {
  username: string;
  password: string;
}

export const signupUser = async ({ username, password }: SignupPayload) => {
  try {
    const data = await axios.post("/user/create", { username, password });
    if (data.status !== 200) {
      throw data;
    }
    return data.data;
  } catch (err: any) {
    if (err?.response?.data?.status === 200) {
      return err?.response?.data;
    }
    throw err;
  }
};

interface GetUserPayload {
  user_id: string;
}

export const getUser = async ({ user_id }: GetUserPayload) => {
  const user = getUserInfo();

  try {
    const data = await axios.get(`/user/get/${user_id}`, {
      headers: { Authorization: user.token },
    });
    if (data.status !== 200) {
      throw data;
    }
    return data.data;
  } catch (err: any) {
    return err?.response?.data;
  }
};
