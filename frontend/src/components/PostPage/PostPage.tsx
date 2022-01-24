import { useQuery } from "react-query";
import { useParams } from "react-router-dom";
import { Post } from "../Posts/Post";
import { ApiResponse } from "../Posts/Posts";

export const PostPage = () => {
  const { slug } = useParams();
  const { isLoading, data } = useQuery<ApiResponse<Post>>(
    `post-${slug}`,
    async () => {
      const response = await fetch(`http://localhost:8080/posts/${slug}`);
      return response.json();
    },
    { refetchOnWindowFocus: false }
  );

  if (isLoading) {
    return <></>;
  }

  return <Post post={data?.data} isPreview={false}></Post>;
};
