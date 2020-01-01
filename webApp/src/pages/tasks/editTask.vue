<template>
    <div>
        <van-nav-bar title="修改任务"
                     left-text="返回"
                     left-arrow
                     @click-left="goBack"/>
        <van-cell title="基本信息"/>
        <van-cell-group>
            <van-field v-model="taskInfo.title" label="活动标题："/>
            <van-field
                    v-model="taskInfo.expire" center clearable label="结束时间：">
                <van-button slot="button" size="small" type="info" @click="showTimePicker=true">选择</van-button>
            </van-field>
            <van-field v-model="taskInfo.status" label="活动状态："/>
        </van-cell-group>
        <van-cell title="详细步骤"/>
        <van-cell-group>
            <van-field
                    v-for="(s, key) in taskInfo.steps"
                    :key="key"
                    :value="s"/>
        </van-cell-group>
        <van-button type="info" size="mini" @click="taskInfo.steps.push('')">添加</van-button>
        <van-cell title="图片教程"/>
        <van-uploader
                :after-read="afterRead"
                multiple
                v-model="tmpImages"
                :max-count="10"/>
        <van-cell title="评论内容"/>
        <van-uploader
                :after-read="afterRead"
                multiple
                v-model="tmpImages"
                :max-count="10"/>
        <van-field v-model="taskInfo.commentBase" label="评论内容："/>
        <van-popup
                title="请选择"
                v-model="showTimePicker"
                position="bottom"
                :style="{ height: '30%' }">
            <van-datetime-picker
                    v-model="currentDate"
                    @cancel="showTimePicker=false"
                    @confirm="confirmDatetime"
                    type="datetime"
                    :min-date="minDate"
                    :max-date="maxDate"
            />
        </van-popup>
    </div>
</template>
<script>
    export default {
        data() {
            return {
                minHour: 10,
                maxHour: 20,
                minDate: new Date(),
                maxDate: new Date(2029, 12, 31),
                currentDate: new Date(),
                showTimePicker: false,
                taskInfo: {
                    title: "评价有礼",
                    status: "进行中",
                    expire: "2020/1/30",
                    description: "1、打开饿了么，点击订单-找到等待评价的订单\n2、返回本页，复制评价，bao保存图片\n3、给商家骑手打好评，在评价里写网页li里的评价，上传网页里的图片，点击提交评价\n4、到商家页面，点击评价，点击最新，截图，回到网页上传，等待客服审核。任务审核完成，佣金直接到你的支付宝钱包",
                    guideImages: [
                        'https://img.yzcdn.cn/vant/apple-1.jpg',
                        'https://img.yzcdn.cn/vant/apple-2.jpg',
                        'https://img.yzcdn.cn/vant/apple-3.jpg',
                        'https://img.yzcdn.cn/vant/apple-4.jpg',
                        'https://img.yzcdn.cn/vant/apple-4.jpg',
                    ],
                    commentBase: "产品不错，快递小哥给力",
                    steps: [""]
                },

                tmpImages: [],
            }
        },
        methods: {
            afterRead: function () {
                console.log(this.tmpImages)
            },
            confirmDatetime:function() {
                console.log(this.currentDate.getTime());
                this.showTimePicker = false
            },
            goBack: function () {
                this.$router.go(-1)
            },
        }
    }
</script>