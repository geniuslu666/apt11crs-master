<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      title="绑定车辆"
      :style="{
        width: dialogWidth,
      }"
    >
      <n-scrollbar style="max-height: 87vh" class="pr-5">
        <n-spin :show="loading" description="请稍候...">
          <n-form
            ref="formRef"
            :model="formValue"
            :label-placement="settingStore.isMobile ? 'top' : 'left'"
            :label-width="100"
            class="py-4"
          >
            <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
              <n-gi span="1" >
                <n-form-item label="选择车辆" path="carId" :show-require-mark="true">
                  <n-select
                    placeholder="请选择车辆"
                    v-model:value="formValue.carId"
                    :options="carList"
                    label-field="carName"
                    value-field="id"
                    clearable
                    filterable
                    style="width: 200px"
                  />
                </n-form-item>
              </n-gi>
            </n-grid>
          </n-form>
        </n-spin>
      </n-scrollbar>
      <template #action>
        <n-space>
          <n-button @click="closeForm">
            取消
          </n-button>
          <n-button type="info" :loading="formBtnLoading" @click="confirmForm">
            确定
          </n-button>
        </n-space>
      </template>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed } from 'vue';
import {BindCar, View} from '@/api/carDriver';
import { useProjectSettingStore } from '@/store/modules/projectSetting';
import { useMessage } from 'naive-ui';
import { adaModalWidth } from '@/utils/hotgo';
import {List} from "@/api/carCar";

const emit = defineEmits(['reloadTable']);
const message = useMessage();
const settingStore = useProjectSettingStore();
const loading = ref(false);
const showModal = ref(false);
const formValue = ref({
  id: 0,
  carId: null,
});
const formRef = ref<any>({});
const formBtnLoading = ref(false);
const dialogWidth = computed(() => {
  return adaModalWidth(840);
});
const selectShow = ref(false);
const carList = ref([]);

function openModal(id,carId) {
  selectShow.value = false
  showModal.value = true;

  // 编辑
  loading.value = true;

  formValue.value.id = id;
  formValue.value.carId = carId ? carId : null;
  console.log("---------formValue--------------",formValue.value)
  List({
    status: 1,
    Pagination: false
  })
    .then((res) => {
      carList.value = res.list;
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
      BindCar(formValue.value).then((_res) => {
        message.success('操作成功');
        formBtnLoading.value = false;
        setTimeout(() => {
          closeForm();
          emit('reloadTable');
        });
      }).catch((err) => {
        formBtnLoading.value = false;
      });
    } else {
      formBtnLoading.value = false;
      message.error('请填写完整信息');
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


