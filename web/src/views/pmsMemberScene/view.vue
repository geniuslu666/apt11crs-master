<template>
  <div class="member-admin-page">
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="会员等级场景详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" column="1">
            <n-descriptions-item>
              <template #label>
                场景名称
              </template>
              {{ formValue.sceneName }}
            </n-descriptions-item>
            <n-descriptions-item label="是否允许积分获取">
              <n-switch v-model:value="formValue.isGetOpen" :unchecked-value="2" :checked-value="1" :disabled="true"/>
            </n-descriptions-item>
            <n-descriptions-item label="是否允许积分抵扣">
              <n-switch v-model:value="formValue.isPayOpen" :unchecked-value="2" :checked-value="1" :disabled="true"/>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                限制金额
              </template>
              {{ formValue.limitMoney }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                获得积分的抵扣比例
              </template>
              {{ formValue.getRate }} <span class="c999">%</span>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                能够使用的总金额比例积分
              </template>
              {{ formValue.payRate }} <span class="c999">%</span>
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
import { View } from '@/api/pmsMemberScene';
import { State, newState, options } from './model';
import { adaModalWidth, getOptionLabel, getOptionTag } from '@/utils/hotgo';
import { getFileExt } from '@/utils/urlUtils';

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

