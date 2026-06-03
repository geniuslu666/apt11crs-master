<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="提现申请表详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" :column=1>
            <n-descriptions-item>
              <template #label>
                提现单号
              </template>
              {{ formValue.withdrawSn }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                司机ID
              </template>
              {{ formValue.driverId }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                司机名称
              </template>
              {{ formValue.driverDetail.name }}
            </n-descriptions-item>

            <n-descriptions-item>
              <template #label>
                提现状态
              </template>
              <template v-if="formValue.withdrawStatus == 'WAIT'">
                <span style="color: green">待审核</span>
              </template>
              <template v-else-if="formValue.withdrawStatus == 'SUCCESS'" >
                <span style="color: blue">已审核</span>
              </template>
              <template v-else-if="formValue.withdrawStatus == 'FAIL'">
                <span style="color: orangered">审核失败</span>
              </template>
              <template v-else>
                --
              </template>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                提现金额
              </template>
              {{ formValue.withdrawAmount }} JPY
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                到账金额
              </template>
              {{ formValue.arrivalAmount }} JPY
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                提现手续费比例
              </template>
              {{ formValue.serviceCharge }} %
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
               转账状态
              </template>
              <template v-if="formValue.transfer == 1">
               待转账
              </template>
              <template v-else-if="formValue.transfer == 2">
                已转账
              </template>
              <template v-else>
                --
              </template>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                审核备注
              </template>
              {{ formValue.applyRemark }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                审核时间
              </template>
              {{ formValue.applyAt }}
            </n-descriptions-item>
            <n-descriptions-item v-if="formValue.transfer == 2">
              <template #label>
                转账时间
              </template>
              {{ formValue.updatedAt }}
            </n-descriptions-item>
          </n-descriptions>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
import { computed, ref } from 'vue';
import { useMessage } from 'naive-ui';
import { View } from '@/api/carDriverWithdraw';
import { State, newState } from './model';
import { adaModalWidth } from '@/utils/hotgo';

const message = useMessage();
const loading = ref(false);
const showModal = ref(false);
const formValue = ref(newState(null));
const dialogWidth = computed(() => {
  return adaModalWidth(580);
});
const fileAvatarCSS = computed(() => {
  return {
    '--n-merged-size': `var(--n-avatar-size-override, 80px)`,
    '--n-font-size': `18px`,
  };
});

//下载
function download(url: string) {
  window.open(url);
}

function openModal(state: State) {
  showModal.value = true;
  loading.value = true;
  View({ id: state.id })
    .then((res) => {
      formValue.value = res;
      console.log('formValue.value',formValue.value);
    })
    .finally(() => {
      loading.value = false;
    });
}

defineExpose({
  openModal,
});
</script>

<style lang="less" scoped></style>


