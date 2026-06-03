<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="接送机地点管理详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" column="1">
            <n-descriptions-item>
              <template #label>
                地点类型id
              </template>
              {{ formValue.typeId }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                地点名称（后台）
              </template>
              {{ formValue.name }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                地点名称（app-多语）
              </template>
              {{ formValue.subName }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                机场代码
              </template>
              {{ formValue.airportCode }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                详细地址-多语
              </template>
              {{ formValue.detailAddress }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                谷歌纬度
              </template>
              {{ formValue.ggLat }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                谷歌经度
              </template>
              {{ formValue.ggLng }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                纬度
              </template>
              {{ formValue.lat }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                经度
              </template>
              {{ formValue.lng }}
            </n-descriptions-item>
            <n-descriptions-item label="状态1、启用 2、禁用">
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
import { View } from '@/api/carAddress';
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


