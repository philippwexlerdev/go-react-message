import axios from "../config/axios";
import { Reply } from "../types/Reply";
import { getUserInfo } from "../util/getUserInfo";

interface GetRepliesPayload {
  post_id: string;
}

export const getReplies = async ({ post_id }: GetRepliesPayload) => {
  const user = getUserInfo();

  try {
    const data = await axios.get(`/post-response/get-all/${post_id}`, {
      headers: { Authorization: user.token },
    });

    if (data.status !== 200) {
      throw data;
    }
    return data.data as Reply[];
  } catch (err: any) {
    if (Array.isArray(err.response.data)) {
      return err.response.data as Reply[];
    }
  }
};

interface CreateReply {
  post_id: string;
  user_id: string;
  message: string;
}

export const createReply = async ({
  post_id,
  user_id,
  message,
}: CreateReply) => {
  const user = getUserInfo();

  try {
    await axios.post(
      `/post-response/create`,
      { post_id, user_id, message },
      {
        headers: { Authorization: user.token },
      }
    );

    return;
  } catch (err: any) {
    return;
  }
};

interface UpdateReply {
  id: string;
  message: string;
}

export const updateReply = async ({ id, message }: UpdateReply) => {
  const user = getUserInfo();

  try {
    await axios.put(
      `/post-response/update/${id}`,
      { message },
      {
        headers: { Authorization: user.token },
      }
    );

    return;
  } catch (err: any) {
    return;
  }
};

interface DeleteReply {
  id: string;
}

export const deleteReply = async ({ id }: DeleteReply) => {
  const user = getUserInfo();

  try {
    await axios.delete(`/post-response/delete/${id}`, {
      headers: { Authorization: user.token },
    });

    return;
  } catch (err: any) {
    return;
  }
};
