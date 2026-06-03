<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="客户档案详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" column="1">
            <n-descriptions-item>
              <template #label>
                客户档案ID
              </template>
              {{ formValue.uid }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                名
              </template>
              {{ formValue.firstName }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                姓
              </template>
              {{ formValue.lastName }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                名的假名
              </template>
              {{ formValue.firstNameKana }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                姓的假名
              </template>
              {{ formValue.lastNameKana }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                全名
              </template>
              {{ formValue.fullName }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                语言
              </template>
              {{ formValue.language }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                电子邮件
              </template>
              {{ formValue.email }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                电话
              </template>
              {{ formValue.phone }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                国籍
              </template>
              {{ formValue.nationality }}
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
  import { View } from '@/api/pmsGuestProfile';
  import { State, newState } from './model';
  import { adaModalWidth} from '@/utils/hotgo';

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