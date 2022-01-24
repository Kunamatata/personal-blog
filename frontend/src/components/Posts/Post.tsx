import { Link } from "react-router-dom";
import snarkdown from "snarkdown";

interface PostProps {
  post?: Post;
  isPreview: boolean;
}

export interface Post {
  slug: string;
  title: string;
  content: string;
  createdAt: string;
  updatedAt: string;
  deletedAt: string;
}

const pluralize = (word: string, count: number) => {
  return count === 1 ? word : word + "s";
};

const PREVIEW_WORDS_LENGTH = 200;

export const Post = (props: PostProps) => {
  const { post } = props;
  if (!post) {
    return <div></div>;
  }

  const estimatedReadingTime = Math.ceil(
    post.content.split(" ").length / 200
  ).toFixed(0);
  const wordCount = post.content.split(" ").length;
  const postDate = new Date(post.createdAt).toLocaleString();
  const previewContent = post.content
    .split(" ")
    .slice(0, PREVIEW_WORDS_LENGTH)
    .join(" ");
  return (
    <div className="border-post-border border mb-6 p-6" key={post.slug}>
      <h2 className="text-xl font-bold">{post.title}</h2>
      <h3 className="italic opacity-40">
        {postDate} - {wordCount} {pluralize("word", wordCount)} - Estimated
        reading time {estimatedReadingTime}{" "}
        {pluralize("minute", Number(estimatedReadingTime))}
      </h3>
      {previewContent.split(" ").length < wordCount ? (
        <article
          className="prose prose-invert"
          dangerouslySetInnerHTML={{
            __html: snarkdown(
              props.isPreview ? `${previewContent}...` : post.content
            ),
          }}
        ></article>
      ) : (
        <article
          className="prose prose-invert"
          dangerouslySetInnerHTML={{ __html: snarkdown(post.content) }}
        ></article>
      )}
      <div className="pt-6 opacity-40 italic">
        {props.isPreview ? (
          <Link to={`/posts/${post.slug}`}>Read more...</Link>
        ) : (
          ""
        )}
      </div>
    </div>
  );
};
