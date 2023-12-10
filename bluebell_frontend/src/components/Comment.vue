<template>
  <div>
    <!-- 评论区 -->
    <div class="comment-container">
      <!-- 评论列表 -->
      <div v-if="comments.length > 0" class="comment-list">
        <div v-for="comment in comments" :key="comment.comment_id" class="comment">
          <div class="c-left">
            <div class="line"></div>
            <div class="c-arrow"></div>
          </div>
          <div class="c-right">
            <div class="c-user-info">
              <div class="name">{{ comment.author_name }}</div>
              <div class="num, el-icon-date">{{formatCreateTime(comment.create_time)}}</div>
            </div>
            <div class="c-content">{{ comment.content }}</div>
          </div>
        </div>
      </div>
      <div v-else> 暂无评论 </div>

      <!-- 评论输入框 -->
      <div class="comment-input">
        <textarea v-model="newComment.content" placeholder="在这里输入评论"></textarea>
        <button @click="submitComment">发表评论</button>
      </div>
    </div>
  </div>
</template>

<script>
import Vue from 'vue';

export default {
  name: 'Comment',
  props: {
    sourceId: {
      type: Number,
      required: false,
      default: -1,
    },
  },
  data() {
    return {
      comments: [], // 存放评论列表
      newComment: {
        content: '',
      }, // 存放新评论
    };
  },
  methods: {
    // 格式化时间
    formatCreateTime(createTime) {
      const dateObject = new Date(createTime);
      const year = dateObject.getFullYear();
      const month = (dateObject.getMonth() + 1).toString().padStart(2, '0');
      const day = dateObject.getDate().toString().padStart(2, '0');
      const hours = dateObject.getHours().toString().padStart(2, '0');
      const minutes = dateObject.getMinutes().toString().padStart(2, '0');
      const seconds = dateObject.getSeconds().toString().padStart(2, '0');

      return `${year}/${month}/${day},${hours}:${minutes}:${seconds}`;
    },
    // 获取评论列表
    getComments() {
      this.$axios({
        method: 'get',
        url: '/comment',
        params: {
          // ids: [this.sourceId],
          ids: this.sourceId,
        },
      })
          .then((response) => {
            if (response.code === 1000) {
              this.comments = response.data;
            } else {
              console.error(response.msg);
            }
          })
          .catch((error) => {
            console.error(error);
          });
    },
    // 提交评论
    submitComment() {
      if (!this.newComment.content.trim()) {
        Vue.prototype.$message.error('评论内容不能为空');
        return;
      }

      this.$axios({
        method: 'post',
        url: '/comment',
        data: {
          post_id: this.sourceId,
          content: this.newComment.content,
        },
      })
          .then((response) => {
            if (response.code === 1000) {
              Vue.prototype.$message.success('评论成功');
              // 清空输入框
              this.newComment.content = '';
              // 刷新评论列表
              this.getComments();
            } else {
              Vue.prototype.$message.error(response.msg);
            }
          })
          .catch((error) => {
            console.error(error);
          });
    },
  },
  mounted() {
    // 获取评论列表
    this.getComments();
  },
};
</script>

<style lang="less" scoped>
.comment-container {
  margin: 32px;
  margin-left: 12px;
  padding-bottom: 30px;
  position: relative;
  border: 1px solid #edeff1;
  border-radius: 4px;
}

.comment-list {
  width: 100%;
  height: auto;
  position: relative;
  border-bottom: 1px solid rgba(0, 0, 0, 0.2);
}

.comment {
  width: 100%;
  height: auto;
  position: relative;
  display: flex;
  margin-bottom: 30px;
  border-bottom: 1px solid #edeff1;
  padding-bottom: 15px;
}

.c-left {
  position: absolute;
  left: 20px;
}

.c-right {
  margin-left: 40px;
  padding-right: 10px;
}

.c-user-info {
  margin-bottom: 10px;
  font-weight: bold;
}

.name {
  color: #003f6d;
  margin-bottom: 10px;
}

.num {
  color: #7c7c7c;
}

.c-content {
  font-family: Noto Sans, Arial, sans-serif;
  font-size: 14px;
  font-weight: 400;
  line-height: 21px;
  word-break: break-word;
  color: rgb(26, 26, 27);
}

.comment-input {
  width: 100%;
  position: relative;
  margin-top: 30px;
  border-top: 1px solid #edeff1;

  textarea {
    width: 100%;
    height: 100px;
    margin-bottom: 10px;
    padding: 10px;
    border: 1px solid #edeff1;
    border-radius: 4px;
  }

  button {
    width: 40%;
    height: 40px;
    background-color: #003f6d;
    color: #ffffff;
    border-radius: 4px;
    cursor: pointer;
    margin: 0 auto;
    display: block;
  }
}
</style>

