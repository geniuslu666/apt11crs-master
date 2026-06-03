<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="物业详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" column="1">
            <n-descriptions-item>
              <template #label>
                物业名称
              </template>
              {{ formValue.name }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                图标
              </template>
              <n-image style="margin-left: 10px; height: 100px; width: 100px" :src="formValue.icon"/>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                封面
              </template>
              <div class="upload-card"  v-show="formValue.cover !== ''" @click="download(formValue.cover)">
                <div class="upload-card-item" style="height: 100px; width: 100px">
                  <div class="upload-card-item-info">
                    <div class="img-box">
                      <n-avatar :style="fileAvatarCSS">
                        {{ getFileExt(formValue.cover) }}
                      </n-avatar>
                    </div>
                  </div>
                </div>
              </div>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                货币
              </template>
              {{ formValue.currency }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                语言
              </template>
              {{ formValue.language }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                时区
              </template>
              {{ formValue.timeZone }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                地址
              </template>
              {{ formValue.address }}
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
  import { View } from '@/api/pmsProperty';
  import { State, newState } from './model';
  import { adaModalWidth } from '@/utils/hotgo';
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