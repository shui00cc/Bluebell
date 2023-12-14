<template>
  <div>
    <div class="content">
      <div class="left">
        <SideBar></SideBar>
      </div>
      <div class="center-right">
        <p class="site-message">
          <span class="forum-title">"随心记"Golang编程论坛</span>
        </p>
        <p class="main-text">
          此论坛由作者历时一个多月code by code搭建完成，<br>包含常见论坛所需的注册登录、评论等功能，
          同时含有多个频道，<br>可以在这里随时记录你想说的话、发表你对于编程学习的看法。
        </p>
        <span v-if="!showTechStack" class="tech-stack-link" @click="toggleTechStack">点击了解技术栈</span>
        <transition name="fade">
          <div v-if="showTechStack" class="tech-stack">
            <p>前后端及数据库:Go+Gin+Vue+Element-ui+MySQL+Redis</p>
            <p>
              项目包含 Viper配置管理、Zap日志记录、优雅关机、雪花算法生成分布式ID、JWT认证、air热重载、
              Swagger接口文档、令牌桶限流、pprof性能测试等实用功能，想要了解更多请前往 <a href="https://github.com/shui00cc/Bluebell">GitHub</a>
            </p>
          </div>
        </transition>
      </div>
    </div>
  </div>
</template>

<script>
import SideBar from '../components/SideBar.vue';

export default {
  name: "About",
  components: { SideBar },
  data() {
    return {
      showTechStack: false,
    };
  },
  methods: {
    toggleTechStack() {
      this.showTechStack = !this.showTechStack;
    },
  },
};
</script>

<style scoped lang="less">
.content {
  max-width: 100%;
  box-sizing: border-box;
  display: flex;
  flex-direction: row;
  justify-content: center;
  margin: 48px auto 0;
  padding: 20px 24px;
  background: #A7BFE8;
}

.left {
  width: 312px;
  height: fit-content;
  margin-top: 28px;
  margin-right: 24px; /* Add margin to create a gap */
  background: #fff;
  border-radius: 4px;
}

.center-right {
  flex: 1;
  margin-top: 28px;
  padding: 0 24px;
  text-align: justify;
  background: #fff; /* White background */
}

.site-message {
  font-size: 28px; /* Increased font size */
  color: #000; /* Black text color */
  margin-bottom: 18px; /* Added margin to separate from the main text */
  margin-top: 18px; /* Added margin to separate from the main text */
}

.forum-title {
  font-weight: bold; /* Bold font for the title */
}

.main-text {
  font-size: 18px; /* Increased font size for the main text */
  line-height: 2.0; /* Increased line height for better readability */
}

.tech-stack {
  font-size: 18px; /* Increased font size for the main text */
  line-height: 2.0; /* Increased line height for better readability */
  margin-top: 20px;
}

.tech-stack-link {
  font-size: 18px; /* Increased font size for the main text */
  line-height: 2.0; /* Increased line height for better readability */
  cursor: pointer;
  color: #00a6ff; /* Blue link color */
}

.fade-enter-active, .fade-leave-active {
  transition: opacity 0.5s;
}

.fade-enter, .fade-leave-to {
  opacity: 0;
}
</style>
