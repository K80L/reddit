<script>
  let { post } = $props();
  $inspect(post);
</script>

<div class="card">
  <h5>r/{post.Subreddit.Name}</h5>
  <h3>{post.Title}</h3>
  <p>{post.Content}</p>
  <div class="post--actions--container">
    <div class="post--actions">
      <button>+</button>
      <span>{post.Likes?.length || 0}</span>
      <button>-</button>
    </div>
    <div class="post--actions">
      <button><img src="/icons8-comment-16.svg" alt="comment" /></button>
    </div>
  </div>
</div>

<style>
  .card {
    border-bottom: 1px solid #27282a;
    min-height: 100px;
    width: 100%;
    padding: 0.25rem 1rem;
  }
  .post--actions--container {
    display: flex;
    gap: 0.75rem;
  }
  .post--actions {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    background: #333d42;
    border-radius: 15px;
    button {
      background: none;
      border: none;
      color: #fff;
    }
  }
</style>
