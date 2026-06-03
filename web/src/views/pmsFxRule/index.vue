<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-card
        :bordered="false"
        class="proCard"
        :header-style="{
                    padding: '25px 20px 20px',
                  }"
        :content-style="{
                    padding: '0 20px 20px',
                  }"
      >
        <template #header>
          <text style="font-weight: 500;font-size: 20px;color: #3D3D3D;line-height: 28px;">分销规则</text>
        </template>
        <div class="statusTab">
          <div :class="tab == 0 ? 'active' : ''"><span @click="handleTab(0)">全局规则</span></div>
          <div :class="tab == 1 ? 'active' : ''"><span @click="handleTab(1)">邀新奖励</span></div>
          <div :class="tab == 2 ? 'active' : ''"><span @click="handleTab(2)">下单奖励</span></div>
        </div>
        <GlobalRule ref="globalRuleRef" v-if="tab == 0"/>
        <InviteReward ref="inviteRewardRef" v-if="tab == 1"/>
        <OrderReward ref="orderRewardRef" v-if="tab == 2"/>
      </n-card>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
import GlobalRule from "@/views/pmsFxRule/global_rule.vue";
import InviteReward from "@/views/pmsFxRule/invite_reward.vue";
import OrderReward from "@/views/pmsFxRule/order_reward.vue";
import {ref} from "vue";

const show = ref(false);
const tab = ref(0);
const globalRuleRef = ref();
const inviteRewardRef = ref();
const orderRewardRef = ref();

function handleTab(e){
  tab.value = e;
}
</script>
<style lang="less" scoped>
.statusTab{
  display: flex;
  margin-bottom: 4px;
  div{
    margin-right: 5px;
    padding: 0 16px;
    height: 30px;
    line-height: 30px;
    text-align: center;
    font-size: 14px;
    color: #4E5969;
    span{
      cursor: pointer;
    }
    &.active{
      background: #F2F3F8;
      border-radius: 30px;
      color: #1664FF;
      font-weight: 500;
    }
  }
}
</style>
