<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="门店管理详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" column="1">
            <n-descriptions-item>
              <template #label>
                名称
              </template>
              {{ formValue.name }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                负责人
              </template>
              {{ formValue.headName }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                联系电话
              </template>
              {{ formValue.phone }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                营业时间
              </template>
              {{ formValue.openTime }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                地区省级ID
              </template>
              {{ formValue.areaPid }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                地区市级ID
              </template>
              {{ formValue.areaId }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                详细地址
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
            <n-descriptions-item>
              <template #label>
                创建时间
              </template>
              {{ formValue.createAt }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                更新时间
              </template>
              {{ formValue.updateAt }}
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
import { View } from '@/api/spaStore';
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


