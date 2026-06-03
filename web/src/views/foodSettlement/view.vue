<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="结算模式详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" column="1">
            <n-descriptions-item label="结算方式 ">
              <n-tag :type="getOptionTag(options.settlement_type, formValue?.type)" size="small" class="min-left-space">
                {{ getOptionLabel(options.settlement_type, formValue?.type) }}
              </n-tag>
            </n-descriptions-item>
            <n-descriptions-item label="结算周期">
              <n-tag :type="getOptionTag(options.settlement_citcle, formValue?.cycle)" size="small" class="min-left-space">
                {{ getOptionLabel(options.settlement_citcle, formValue?.cycle) }}
              </n-tag>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                结算比例
              </template>
              {{ formValue.rate }}
            </n-descriptions-item>
            <n-descriptions-item label="结算成本控制">
              <template v-for="(item, key) in formValue?.cost" :key="key">
                <n-tag :type="getOptionTag(options.settlement_cost, item)" size="small" class="min-left-space">
                  {{ getOptionLabel(options.settlement_cost, item) }}
                </n-tag>
              </template>
            </n-descriptions-item>
            <n-descriptions-item label="是否提现审核">
              <n-switch v-model:value="formValue.isAuditWithdraw" :unchecked-value="2" :checked-value="1" :disabled="true"/>
            </n-descriptions-item>
            <n-descriptions-item label="状态">
              <n-tag :type="getOptionTag(options.sys_normal_disable, formValue?.status)" size="small" class="min-left-space">
                {{ getOptionLabel(options.sys_normal_disable, formValue?.status) }}
              </n-tag>
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
import { View } from '@/api/foodSettlement';
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


