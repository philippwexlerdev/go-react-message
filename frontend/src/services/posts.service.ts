import axios from "../config/axios";
import { Post } from "../types/Post";
import { getUserInfo } from "../util/getUserInfo";

export const getPosts = async () => {
  const user = getUserInfo();

  try {
    const data = await axios.get("/post/get-all", {
      headers: { Authorization: user.token },
    });
    if (data.status !== 200) {
      throw data;
    }
    return data.data as Post[];
  } catch (err: any) {
    if (Array.isArray(err.response.data)) {
      return err.response.data as Post[];
    }
    throw err;
  }
};

interface CreatePostPayload {
  message: string;
}

export const createPost = async ({ message }: CreatePostPayload) => {
  const user = getUserInfo();

  try {
    const data = await axios.post(
      "/post/create",
      { message, user_id: user.user_id },
      {
        headers: { Authorization: user.token },
      }
    );
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

interface DeletePostPayload {
  postId: string;
}

export const deletePost = async ({ postId }: DeletePostPayload) => {
  const user = getUserInfo();

  try {
    await axios.delete(`/post/delete/${postId}`, {
      headers: { Authorization: user.token },
    });
    return;
  } catch (err: any) {
    return;
  }
};

interface EditPostPayload {
  postId: string;
  message: string;
}

export const editPost = async ({ postId, message }: EditPostPayload) => {
  const user = getUserInfo();

  try {
    await axios.put(
      `/post/update/${postId}`,
      {
        message,
      },
      {
        headers: { Authorization: user.token },
      }
    );
    return;
  } catch (err: any) {
    return;
  }
};
