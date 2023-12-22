document.addEventListener('DOMContentLoaded', () => {
    fetch('http://localhost:8080/post')
      .then(response => response.json())
      .then(data => displayPosts(data));
  });
  
  function displayPosts(posts) {
    const postBody = document.getElementById('postBody');
    postBody.innerHTML = '';
  
    posts.forEach(post => {
      const row = document.createElement('tr');
      row.innerHTML = `
        <td>${post.id}</td>
        <td>${post.title}</td>
        <td>${post.createDate}</td>
      `;
      row.addEventListener('click', () => showPostDetails(post.id));
      postBody.appendChild(row);
    });
  }
  
  function showPostDetails(postId) {
    window.location.href = `detail.html?id=${postId}`;
  }
  