<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content :body-content-style="{
                    padding: '20px',
                  }" :footer-style="{
                    padding: '12px 20px',
                  }" :header-style="{
                    padding: '20px',
                  }" closable>
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">房间编辑
          </div>
        </template>
        <template #footer>
          <n-button style="width: 70px;height: 35px;margin-right: 10px" @click="closeForm">
            取消
          </n-button>
          <n-button :loading="formBtnLoading" style="width: 70px;height: 35px;" type="info"
                    @click="confirmForm">
            保存
          </n-button>
        </template>
        <n-spin :show="loading" description="请稍候...">
          <n-form
            ref="formRef"
            :model="formValue"
            :rules="rules"
            label-placement="top"
            label-width="auto"
          >
            <n-grid :cols="1">
              <n-gi>
                <n-form-item label="第三方系统的ID" path="uid">
                  <n-input v-model:value="formValue.uid" placeholder="请输入第三方系统的ID"
                           style="width: 300px"/>
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="房型ID" path="rtUid">
                  <n-input v-model:value="formValue.rtUid" placeholder="请输入房型ID"
                           style="width: 300px"/>
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="房间号" path="roomNo">
                  <n-input v-model:value="formValue.roomNo" placeholder="请输入房间号"
                           style="width: 300px"/>
                </n-form-item>
              </n-gi>
            </n-grid>
          </n-form>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
  import { ref, computed } from 'vue';
  import { Edit, View } from '@/api/pmsRoomUnit';
  import { State, newState, rules } from './model';
  import { useProjectSettingStore } from '@/store/modules/projectSetting';
  import {NButton, useMessage} from 'naive-ui';
  import { adaModalWidth } from '@/utils/hotgo';

  const emit = defineEmits(['reloadTable']);
  const message = useMessage();
  const settingStore = useProjectSettingStore();
  const loading = ref(false);
  const showModal = ref(false);
  const formValue = ref<State>(newState(null));
  const formRef = ref<any>({});
  const formBtnLoading = ref(false);
  const dialogWidth = computed(() => {
    return adaModalWidth(650);
  });

  function openModal(state: State) {
    showModal.value = true;

    // 新增
    if (!state || state.id < 1) {
      formValue.value = newState(state);

      return;
    }

    // 编辑
    loading.value = true;
    View({ id: state.id })
      .then((res) => {
        formValue.value = res;
      })
      .finally(() => {
        loading.value = false;
      });
  }

  function confirmForm(e) {
    e.preventDefault();
    formBtnLoading.value = true;
    formRef.value.validate((errors) => {
      if (!errors) {
        Edit(formValue.value).then((_res) => {
          message.success('操作成功');
          setTimeout(() => {
            closeForm();
            emit('reloadTable');
          });
        });
      } else {
        message.error('请填写完整信息');
      }
      formBtnLoading.value = false;
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
