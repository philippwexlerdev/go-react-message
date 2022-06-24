import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import Paper from "@mui/material/Paper";
import { usePostsQuery } from "../../queries/usePostsQuery";
import CreatePost from "../../components/CreatePost";
import Post from "../../components/Post";

export default function Home() {
  const { data: posts, isLoading } = usePostsQuery();

  return (
    <Container
      sx={{ display: "flex", flexDirection: "column", gap: 2, marginTop: 2 }}
    >
      <Paper sx={{ padding: 2 }}>
        <h1>Posts</h1>
        {posts?.length === 0 && <h6>No Posts!</h6>}
        <Box sx={{ display: "flex", gap: 2, flexWrap: "wrap" }}>
          {posts?.map((post) => (
            <Post post={post} key={post.id} />
          ))}
        </Box>
      </Paper>
      <Paper sx={{ padding: 2 }}>
        <h1>Create Post</h1>
        <CreatePost />
      </Paper>
    </Container>
  );
}
