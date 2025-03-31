<!DOCTYPE html>
<html lang="zh-CN">

<head>
 <meta charset="UTF-8">
 <meta name="viewport" content="width=device-width, initial-scale=1.0">
 <title>后台管理系统</title>
 <script src="https://cdn.tailwindcss.com"></script>
 <script src="https://unpkg.com/feather-icons"></script>
 <script>
  // 全局功能
  function setPageTitle(title) {
   document.getElementById('pageTitle').textContent = title;
   document.title = title + ' - 后台管理系统';
  }

  function logout() {
   if (confirm('确定要退出系统吗？')) {
    localStorage.clear();
    window.location.href = '/logout';
   }
  }
 </script>
 <style>
  /* 自定义滚动条 */
  ::-webkit-scrollbar {
   width: 8px;
  }

  ::-webkit-scrollbar-track {
   background: #f1f5f9;
  }

  ::-webkit-scrollbar-thumb {
   background: #cbd5e1;
   border-radius: 4px;
  }
 </style>
</head>

<body class="bg-gray-100">
 <!-- 侧边栏 -->
 <aside class="bg-gray-800 text-white w-64 fixed h-full p-4 overflow-y-auto">
  <div class="mb-8">
   <h1 class="text-xl font-bold">后台管理系统</h1>
   <p class="text-gray-400 text-sm mt-1">v2.1.0</p>
  </div>
  <nav>
   <ul class="space-y-2">
    <li>
     <a href="index" class="flex items-center p-2 hover:bg-gray-700 rounded">
      <i data-feather="home" class="w-4 h-4 mr-2"></i> 首页概览
     </a>
    </li>
    <li>
     <a href="userList?page=1&page_size=10" class="flex items-center p-2 hover:bg-gray-700 rounded">
      <i data-feather="users" class="w-4 h-4 mr-2"></i> 用户管理
     </a>
    </li>
    <!-- 更多菜单项... -->
   </ul>
  </nav>
 </aside>

 <!-- 全局顶部导航 -->
 <header class="ml-64 fixed w-[calc(100%-16rem)] bg-white shadow-sm z-10">
  <div class="flex justify-between items-center px-8 py-4">
   <h2 class="text-xl font-bold" id="pageTitle"></h2>
   <div class="flex items-center gap-4">
    <button class="p-2 hover:bg-gray-100 rounded-full">
     <i data-feather="bell"></i>
    </button>
    <div class="flex items-center gap-2">
     <img src="{{.Image}}" class="w-8 h-8 rounded-full">
     <button onclick="logout()" class="text-red-600 hover:text-red-700 flex items-center gap-1">
      <i data-feather="log-out" class="w-5 h-5"></i>
      <span class="hidden sm:inline">退出系统</span>
     </button>
    </div>
   </div>
  </div>
 </header>

 <!-- 主内容容器 -->
 <main class="ml-64 pt-20 p-8" id="mainContent">
  <!-- 继承通用布局 -->
  <script>
   setPageTitle('用户管理');
  </script>

  <div class="space-y-6">
   <!-- 操作栏 -->
   <div class="flex flex-col sm:flex-row justify-between gap-4">
    <div class="flex flex-col sm:flex-row gap-4">
     <!-- 搜索框宽度优化 -->
     <div class="relative flex-[2_2_0%] min-w-[200px]">
      <input type="text" id="searchUsername" placeholder="搜索用户..."
       class="w-full pl-10 pr-4 py-2 border rounded-lg focus:ring-2 focus:ring-blue-500">
      <i data-feather="search" class="absolute left-3 top-2.5 text-gray-400"></i>
     </div>
     <button onclick="searchUsers()" class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700">
      搜索
     </button>

     <!-- 新增状态筛选 -->
     <select class="px-4 py-2 border rounded-lg focus:ring-2 focus:ring-blue-500 flex-1" id="statusFilter" onchange="applyFilter()">
      <option value="">全部状态</option>
      <option value="1" {{if eq .Pagination.CurrentFilter 1}}selected{{end}}>启用</option>
      <option value="0" {{if eq .Pagination.CurrentFilter 0}}selected{{end}}>禁用</option>
     </select>
    </div>
    <button onclick="openUserModal()"
     class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 flex items-center gap-2 whitespace-nowrap">
     <i data-feather="plus"></i>
     新建用户
    </button>
   </div>

   <!-- 用户表格 -->
   <div class="bg-white rounded-xl shadow overflow-x-auto">
    <table class="w-full">
     <thead class="bg-gray-50">
      <tr>
       <th class="px-6 py-3 text-left text-sm font-medium text-gray-500">用户名</th>
       <th class="px-6 py-3 text-left text-sm font-medium text-gray-500">角色</th>
       <th class="px-6 py-3 text-left text-sm font-medium text-gray-500">最后登录</th>
       <th class="px-6 py-3 text-left text-sm font-medium text-gray-500">状态</th>
       <th class="px-6 py-3 text-left text-sm font-medium text-gray-500">操作</th>
      </tr>
     </thead>
     <tbody class="divide-y divide-gray-200">
     {{range .Pagination.UserList}}
      <tr class="hover:bg-gray-50">
       <td class="px-6 py-4">
        <div class="flex items-center gap-3">
         <img src="{{.Image}}" class="w-8 h-8 rounded-full">
         <div>
          <p class="font-medium">{{.Username}}</p>
          <p class="text-sm text-gray-500">{{.Email}}</p>
         </div>
        </div>
       </td>
       <td class="px-6 py-4">
        <span class="px-2 py-1 bg-blue-100 text-blue-800 rounded-full text-sm">{{.Role}}</span>
       </td>
       <td class="px-6 py-4">{{.LastLogin}}</td>
       <td class="px-6 py-4">
        {{if eq .Status 1}}
        <span class="px-2 py-1 bg-green-100 text-green-800 rounded-full text-sm">启用</span>
        {{else}}
        <span class="px-2 py-1 bg-gray-100 text-gray-800 rounded-full text-sm">禁用</span>
        {{end}}
       </td>
       <td class="px-6 py-4">
        <div class="flex gap-2">
         <button class="p-2 hover:bg-gray-100 rounded" onclick="openEditUserModal({{.Username}})">
          <i data-feather="edit" class="w-4 h-4 text-blue-600"></i>
         </button>
         <button class="p-2 hover:bg-gray-100 rounded" onclick="deleteUser({{.Username}})">
          <i data-feather="trash-2"  class="w-4 h-4 text-red-600"></i>
         </button>
        </div>
       </td>
      </tr>
     {{end}}
      <!-- 更多用户数据... -->
     </tbody>
    </table>
   </div>

   <!-- 分页 -->
   <div class="flex justify-between items-center px-4 py-3 bg-white rounded-xl shadow">
    <span class="text-sm text-gray-600">显示 1-{{.Pagination.PerPage}} 页，共 {{.Pagination.Total}} 项</span>
    <div class="flex gap-2">
     <button class="px-3 py-1 rounded hover:bg-gray-100" {{if not .Pagination.HasPrev}}disabled{{end}} onclick="navigateToPage({{.Pagination.PrevPage}},{{.Pagination.PageSize}},{{.Pagination.CurrentFilter}})">上一页</button>
     {{ range .Pagination.PageList }}
     <button class="px-3 py-1 {{if .Active}}bg-blue-100{{end}} rounded hover:bg-gray-100" onclick="navigateToPage({{.Number}},{{$.Pagination.PageSize}},{{$.Pagination.CurrentFilter}})">{{.Number}}</button>
     {{end}}
     <button class="px-3 py-1 rounded hover:bg-gray-100" {{if not .Pagination.HasNext}}disabled{{end}} onclick="navigateToPage({{.Pagination.NextPage}},{{.Pagination.PageSize}},{{.Pagination.CurrentFilter}})">下一页</button>
    </div>
   </div>
  </div>

  <!-- 增加用户模态框 -->
  <div id="userModal" class="hidden fixed inset-0 bg-black/50 flex items-center justify-center p-4">
   <div class="bg-white rounded-xl p-6 w-full max-w-md">
    <h3 class="text-xl font-bold mb-4">新建用户</h3>
    <form class="space-y-4" id="userModel">
     <div>
      <label class="block text-sm font-medium mb-1">用户头像</label>
      <input type="file" accept="image/*" class="w-full px-4 py-2 border rounded-lg focus:ring-2 focus:ring-blue-500" name="avatar" id="avatarUpload">
      <div class="mt-2">
       <img id="avatarPreview" class="w-20 h-20 rounded-full object-cover hidden" alt="预览图片">
      </div>
     </div>
     <div>
      <label class="block text-sm font-medium mb-1">用户名</label>
      <input type="text" class="w-full px-4 py-2 border rounded-lg focus:ring-2 focus:ring-blue-500" name="username">
     </div>
     <div>
      <label class="block text-sm font-medium mb-1">密码</label>
      <input type="password" class="w-full px-4 py-2 border rounded-lg focus:ring-2 focus:ring-blue-500" name="password">
     </div>
     <div>
      <label class="block text-sm font-medium mb-1">邮箱</label>
      <input type="text" class="w-full px-4 py-2 border rounded-lg focus:ring-2 focus:ring-blue-500" name="email">
     </div>
     <div class="mb-4">
      <label class="block text-sm font-medium mb-1">用户状态</label>
      <select class="w-full px-4 py-2 border rounded-lg focus:ring-2 focus:ring-blue-500" id="userStatus" name="status">
       <option value="1">启用</option>
       <option value="0">禁用</option>
      </select>
     </div>
     <div class="flex justify-end gap-2">
      <button type="button" onclick="closeUserModal()" class="px-4 py-2 text-gray-600 hover:bg-gray-100 rounded-lg">
       取消
      </button>
      <button type="submit" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
       提交
      </button>
     </div>
    </form>
   </div>
  </div>

  <!-- 编辑用户模态框 -->
  <div id="editUserModal" class="hidden fixed inset-0 bg-black/50 flex items-center justify-center p-4">
   <div class="bg-white rounded-xl p-6 w-full max-w-md">
    <h3 class="text-xl font-bold mb-4">编辑用户</h3>
    <form class="space-y-4" id="editUserForm">
     <input type="hidden" name="userId" id="userId">
     <input type="hidden" name="currentAvatar" id="currentAvatarPath"> <!-- 添加隐藏输入框 -->
     <div>
      <label class="block text-sm font-medium mb-1">用户头像</label>
      <div class="flex flex-col gap-2">
       <img id="currentAvatar" class="w-20 h-20 rounded-full object-cover mb-2" alt="当前头像">
       <input type="file" accept="image/*" class="w-full px-4 py-2 border rounded-lg focus:ring-2 focus:ring-blue-500" name="avatar" id="editAvatarUpload">
       <div class="mt-2">
        <img id="editAvatarPreview" class="w-20 h-20 rounded-full object-cover hidden" alt="预览图片">
       </div>
      </div>
     </div>
     <div>
      <label class="block text-sm font-medium mb-1">用户名</label>
      <input type="text" class="w-full px-4 py-2 border rounded-lg focus:ring-2 focus:ring-blue-500" name="username" id="editUsername">
     </div>
     <div>
      <label class="block text-sm font-medium mb-1">密码</label>
      <input type="password" class="w-full px-4 py-2 border rounded-lg focus:ring-2 focus:ring-blue-500" name="password" id="editPassword">
     </div>
     <div>
      <label class="block text-sm font-medium mb-1">邮箱</label>
      <input type="text" class="w-full px-4 py-2 border rounded-lg focus:ring-2 focus:ring-blue-500" name="email" id="editEmail">
     </div>
     <div class="mb-4">
      <label class="block text-sm font-medium mb-1">用户状态</label>
      <select class="w-full px-4 py-2 border rounded-lg focus:ring-2 focus:ring-blue-500" name="status" id="editUserStatus">
       <option value="1">启用</option>
       <option value="0">禁用</option>
      </select>
     </div>
     <div class="flex justify-end gap-2">
      <button type="button" onclick="closeEditUserModal()" class="px-4 py-2 text-gray-600 hover:bg-gray-100 rounded-lg">
       取消
      </button>
      <button type="submit" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
       提交
      </button>
     </div>
    </form>
   </div>
  </div>

  <script>
   // 用户管理页面专属逻辑
   function openUserModal() {
    document.getElementById('userModal').classList.remove('hidden');
   }

   function closeUserModal() {
    document.getElementById('userModal').classList.add('hidden');
   }

   // 点击模态框外部关闭
   document.getElementById('userModal').addEventListener('click', (e) => {
    if (e.target === document.getElementById('userModal')) {
     closeUserModal();
    }
   });

   document.getElementById('avatarUpload').addEventListener('change', function(event) {
    const file = event.target.files[0];
    if (file) {
     const reader = new FileReader();
     reader.onload = function(e) {
      const preview = document.getElementById('avatarPreview');
      preview.src = e.target.result;
      preview.classList.remove('hidden');
     };
     reader.readAsDataURL(file);
    }
   });

   // 提交增加表单
   document.getElementById('userModal').addEventListener('submit', (e) => {
    e.preventDefault();
    const formData = new FormData(e.target);
    const data = Object.fromEntries(formData.entries());
    // 整型处理
    data.status = parseInt(data.status, 10);

    fetch(`/register`, {
     method: 'PUT',
     body: formData, // 使用 formData 代替 JSON
    }).then(response => response.json())
            .then(data => {
             alert(data.msg);
             if (data.status === 200) {
              window.location.href = `/userList?page=1&page_size=10`;
             }
            })
            .catch(error => {
             console.error('增加用户信息时出错:', error);
             alert('增加用户信息时出错，请重试。');
            });
   });

   document.getElementById('editAvatarUpload').addEventListener('change', function(event) {
    const file = event.target.files[0];
    if (file) {
     const reader = new FileReader();
     reader.onload = function(e) {
      const preview = document.getElementById('editAvatarPreview');
      preview.src = e.target.result;
      preview.classList.remove('hidden');
     };
     reader.readAsDataURL(file);
    }
   });

   function openEditUserModal(username) {
    // 获取用户信息并填充表单
    fetch(`/edit?username=${username}`)
            .then(response => response.json())
            .then(user => {
             document.getElementById('userId').value = user.data.id;
             document.getElementById('editUsername').value = user.data.username;
             document.getElementById('editPassword').value = user.data.password;
             document.getElementById('editEmail').value = user.data.email;
             document.getElementById('editUserStatus').value = user.data.status;
             // 设置当前头像
             const currentAvatar = document.getElementById('currentAvatar');
             currentAvatar.src = user.data.image; // 假设服务器返回的图片路径是 user.data.image
             // 设置隐藏输入框的值
             document.getElementById('currentAvatarPath').value = user.data.image;
             document.getElementById('editUserModal').classList.remove('hidden');
            });

   }

   function closeEditUserModal() {
    document.getElementById('editUserModal').classList.add('hidden');
   }

   // 点击模态框外部关闭
   document.getElementById('editUserModal').addEventListener('click', (e) => {
    if (e.target === document.getElementById('editUserModal')) {
     closeEditUserModal();
    }
   });

   document.getElementById('editUserForm').addEventListener('submit', (e) => {
    e.preventDefault();
    const formData = new FormData(e.target);
    const data = Object.fromEntries(formData.entries());
    // 整型处理
    data.status = parseInt(data.status, 10);

    fetch(`/edit/${formData.get('userId')}`, {
     method: 'PUT',
     body: formData, // 使用 formData 代替 JSON
    }).then(response => response.json())
            .then(data => {
             alert(data.msg);
             if (data.status === 200) {
              window.location.href = '/userList?page=1&page_size=10';
             }
            })
            .catch(error => {
             console.error('修改用户信息时出错:', error);
             alert('修改用户信息时出错，请重试。');
            });
   });

   function navigateToPage(page,page_size,statusFilter) {
    // 构建基础URL
    let url = `/userList?page=${page}&page_size=${page_size}`;

    // 如果有状态过滤，则添加到URL
    if (statusFilter !== undefined && statusFilter !== "") {
     url += `&status=${statusFilter}`;
    }

    // 跳转到新URL
    window.location.href = url;
   }

   function deleteUser(username) {
    if (confirm('确定要删除该用户吗？')) {
     window.location.href = `/delete?username=${username}`;
    }
   }

   // 搜索用户
   function searchUsers() {
    const searchUsername = document.getElementById('searchUsername').value;
    const urlObject = new URL(`/userList?page=1&page_size=10`,window.location.origin);
    if (searchUsername == ""){
     window.location.href = urlObject.toString();
      return;
    }
    urlObject.searchParams.set('username', searchUsername);
    window.location.href = urlObject.toString();
   }

   function applyFilter() {
    const statusFilter = document.getElementById('statusFilter').value;
    // 拼接上当前域名
    const urlObject = new URL(`/userList?page=1&page_size=10`,window.location.origin);
    if (statusFilter == ""){
     window.location.href = urlObject.toString();
     return;
    }
    urlObject.searchParams.set('status', statusFilter);
    window.location.href = urlObject.toString();
   }
  </script>
 </main>
<script>// 初始化图标
 feather.replace();
</script>
</body>

</html>