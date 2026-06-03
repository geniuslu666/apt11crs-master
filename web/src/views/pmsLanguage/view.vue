<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="语言字典详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" column="1">
            <n-descriptions-item>
              <template #label>
                标签ID
              </template>
              {{ formValue.uuid }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                语言标签
              </template>
              {{ formValue.tag }}
            </n-descriptions-item>
            <n-descriptions-item label="类型">
              <n-tag :type="getOptionTag(options.type, formValue?.type)" size="small" class="min-left-space">
                {{ getOptionLabel(options.type, formValue?.type) }}
              </n-tag>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                标识
              </template>
              {{ formValue.key }}
            </n-descriptions-item>
            <n-descriptions-item label="语言">
              <n-tag :type="getOptionTag(options.language, formValue?.language)" size="small" class="min-left-space">
                {{ getOptionLabel(options.language, formValue?.language) }}
              </n-tag>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                内容
              </template>
              <span v-html="formValue.content"></span>
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
  import { View } from '@/api/pmsLanguage';
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