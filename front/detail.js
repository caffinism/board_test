document.addEventListener('DOMContentLoaded', () => {
    const urlParams = new URLSearchParams(window.location.search);
    const postId = urlParams.get('id');
  
    if (postId) {
      fetch(`http://localhost:8080/post?id=${postId}`)
        .then(response => response.json())
        .then(data => displayPostDetails(data));
    }
  });
  
  function displayPostDetails(post) {
    document.getElementById('postId').textContent = post.id;
    document.getElementById('postTitle').textContent = post.title;
    document.getElementById('postContent').textContent = post.content;
    document.getElementById('postCreateDate').textContent = post.createDate;
    document.getElementById('postUpdateDate').textContent = post.updateDate;
  }
  