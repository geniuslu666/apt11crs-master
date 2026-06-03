<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth" @mask-click="closeForm">
      <n-drawer-content :title="translang('经纬度设置')" closable :native-scrollbar="false">
        <n-spin :show="loading" description="请稍候...">
          <n-form
            ref="formRef"
            :rules="rules"
            :model="formValue"
            label-placement="top"
            label-width="auto"
            class="py-4"
          >
            <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
              <n-gi span="1">
                <n-form-item label="纬度" path="ggLat">
                  <n-input
                    placeholder="请输入纬度"
                    v-model:value="formValue.ggLat"
                    style="width: 250px"
                  />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="经度" path="ggLng">
                  <n-input
                    placeholder="请输入经度"
                    v-model:value="formValue.ggLng"
                    style="width: 250px"
                  />
                </n-form-item>
              </n-gi>
            </n-grid>
          </n-form>
        </n-spin>
        <n-space>
          <n-button @click="closeForm"> 取消 </n-button>
          <n-button type="info" :loading="formBtnLoading" @click="confirmForm"> 确定 </n-button>
        </n-space>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
  import { computed, ref } from 'vue';
  import { Edit } from '@/api/pmsProperty';
  import {useMessage} from 'naive-ui';
  import { translang } from '@/utils/smjcomm';
  import { adaModalWidth } from '@/utils/hotgo';

  const emit = defineEmits(['reloadTable']);
  const message = useMessage();
  const loading = ref(false);
  const showModal = ref(false);
  const dialogWidth = computed(() => {
    return adaModalWidth(620);
  });
  const rules = ref({
    ggLat: {
      required: true,
      trigger: ['blur', 'input'],
      type: 'string',
      message: '请输入纬度',
    },
    ggLng: {
      required: true,
      trigger: ['blur', 'input'],
      type: 'string',
      message: '请输入经度',
    },
  });
  const formValue = ref({
    id: 0,
    ggLat: '',
    ggLng: '',
  });
  const formRef = ref<any>({});
  const formBtnLoading = ref(false);

  function openModal(data) {
    
    loading.value = true;
    formValue.value.id = data.id;
    formValue.value.ggLat = data.ggLat;
    formValue.value.ggLng = data.ggLng;
    loading.value = false;
    showModal.value = true;
  }

  function confirmForm(e) {
    e.preventDefault();
    formBtnLoading.value = true;
    formRef.value.validate((errors) => {
      if (!errors) {
        Edit(formValue.value)
          .then((_res) => {
            message.success('操作成功');
            setTimeout(() => {
              closeForm();
              formBtnLoading.value = false;
              emit('reloadTable');
            });
          })
          .catch((err) => {
            formBtnLoading.value = false;
          });
      } else {
        message.error('验证错误');
        formBtnLoading.value = false;
      }
    });
  }

  function closeForm() {
    showModal.value = false;
    loading.value = false;
  }

  defineExpose({
    openModal,
  });
</script>

<style lang="less"></style>
