<template>
  <div>
    <div class="n-layout-page-header">
<!--      <n-card :bordered="false" title="渠道分销">
      </n-card>-->
    </div>

    <n-card :bordered="false" class="proCard">
      <n-tabs
        type="card"
        class="card-tabs"
        animated
        @onBeforeLeave="handleBeforeLeave"
        :pane-style="{
                    padding: '0',
                  }"
      >
        <n-tab-pane name="channelIndex" tab="优惠券列表"> <IndexList /></n-tab-pane>
        <n-tab-pane name="channelWithdraw" tab="会员优惠券"> <MemberCouponList /> </n-tab-pane>
      </n-tabs>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
  import { onMounted, ref } from 'vue';
  import IndexList from '@/views/pmsCouponType/index.vue';
  import MemberCouponList from '@/views/pmsCouponType/member_coupon.vue';

  const defaultTab = ref('');

  onMounted(() => {
  });

  function handleBeforeLeave(tabName: string) {
    defaultTab.value = tabName;
  }
</script>

<style lang="less" scoped></style>
