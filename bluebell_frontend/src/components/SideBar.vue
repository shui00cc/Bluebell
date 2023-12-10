<template>
	<div class="side-bar-container">
		<div class="title">
			<el-avatar><span><img src="../assets/images/myHead.jpg"></span></el-avatar>
			<el-divider><i class="el-icon-user-solid"></i></el-divider>
		</div>
		<p align="center">
			Visitor count<br>
			<img src="https://profile-counter.glitch.me/shui00cc/count.svg" />
		</p>
		<el-carousel indicator-position="outside">
			<el-carousel-item v-for="item in imgOptions" :key="item.id">
				<img style="object-fit: cover;width:100%;height: 100%;;" :src="item.path" alt="图片" width="300px" />
			</el-carousel-item>
		</el-carousel>

		<div class="message">
			<ul>
				<li>
					<el-tag type="success">
						<i class="el-icon-grape"></i>
            电话: 18223786958
					</el-tag>
				</li>
				<li>
					<el-tag type="warning">
						<i class="el-icon-ice-tea"></i>
						邮箱: 384354037@qq.com
					</el-tag>
				</li>
				<li>
					<el-tag>
						<i class="el-icon-watermelon"></i>
						QQ交流: 384354037
					</el-tag>
				</li>
			</ul>
		</div>
		<div class="mine">
			<el-card class="box-card">
				<div slot="header" class="clearfix" style="text-align:center;">
					<span>作者留言</span>
				</div>
				<!-- v-for="o in 1" :key="o"  -->
				<div class="text item">
          Don't beg, don't chase, and don't be needy, just relax
				</div>
			</el-card>
		</div>
	</div>
</template>
  
<script>
export default {
	name: "SideBar",
	data() {
		return {
			imgOptions: [
				{
					id: 1,
					path: require("@/assets/images/myPhoto1.jpg"),
				},
				{
					id: 2,
					path: require("@/assets/images/myPhoto2.jpg"),
				},
				{
					id: 3,
					path: require("@/assets/images/myPhoto3.jpg"),
				}
			],
			value: new Date(),
			value2: null,
			colors: ["#99A9BF", "#F7BA2A", "#FF9900"],
		};
	},
};
</script>
  
<style type="scss" scoped>
.side-bar-container {
	display: flex;
	flex-direction: column;
	font-size: 14px;
}

.title {
	text-align: center;
	margin-top: 10px;
}


.title .leave-message i {
	font-size: 1.1rem;
}

.title /deep/ .el-avatar img {
	width: 100%;
	height: 100%;
}

.el-carousel {
	width: 100%;
	margin: 0px 0px 10px 0px;
}

.el-carousel img {
	/* width:250px; */
	height: 180px;
}

::v-deep .el-carousel__container {
	height: 200px;
}

.block {
	width: 210px;
	margin: 0 auto;
	text-align: center;
	background-color: rgb(220 169 31 / 29%);
	box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}

.message {
	display: block;
}

.message ul {
	margin: 20px 0px 20px 0px;
	padding: 0px;
}

.message ul li {
	width: 210px;
	height: 32px;
	list-style: none;
	margin-top: 10px;
	margin-left: 70px;
	margin: 0 auto;
	display: block;
	text-align: center;

}

.message ul li .el-tag {
	width: 210px;
	height: 32px;
}

.mine {
	width: 210px;
	margin: 0 auto;

	.box-card {}
}

.mine span {
	text-align: center;
	font-weight: 600;
	color: #58666e;
}

.mine .box-card {
	margin-bottom: 1rem;
}

.text {
	text-align: center;
}

</style>