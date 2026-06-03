<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-card :bordered="false" :header-style="{
                    padding: '20px',
                  }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">小票打印设置</text>
        </template>
        <n-form
          ref="formRef"
          :model="formValue"
          :rules="rules"
          :label-placement="settingStore.isMobile ? 'top' : 'left'"
          :label-width="150"
          class="py-4"
          style="padding-top: 0"
        >
          <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
            <n-gi span="1">
              <n-form-item label="是否开启小票打印" path="isOpen">
                <n-radio-group v-model:value="formValue.isOpen" name="isOpen">
                  <n-space>
                    <n-radio :value="1">
                      开启
                    </n-radio>
                    <n-radio :value="2">
                      关闭
                    </n-radio>
                  </n-space>
                </n-radio-group>
              </n-form-item>
            </n-gi>
            <n-gi span="1" v-if="formValue.isOpen == 1">
              <n-form-item label="打印机" path="group" style="width: 800px">
                <n-select
                  placeholder="请选择打印机"
                  v-model:value="printerIdsArr"
                  :options="printerList"
                  label-field="printerName"
                  value-field="id"
                  clearable
                  filterable
                  multiple
                  style="width: 300px"
                />
              </n-form-item>
            </n-gi>
            <n-gi span="1" v-if="formValue.isOpen == 1">
              <n-form-item label="打印联数" path="printerNum">
                <n-input-group>
                  <n-input-number placeholder="请输入" :show-button="false" :min="0" :precision="0" v-model:value="formValue.printerNum" style="width: 100px" />
                  <n-input-group-label>联</n-input-group-label>
                </n-input-group>
              </n-form-item>
            </n-gi>
          </n-grid>
          <div style="text-align: center">
            <n-space justify="center">
              <n-button type="info" :loading="formBtnLoading" @click="formSubmit">
                确定
              </n-button>
            </n-space>
          </div>
        </n-form>
      </n-card>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
import {ref, onMounted, reactive} from 'vue';
  import { useMessage } from 'naive-ui';
  import { getConfig, updateConfig } from '@/api/sys/config';
  import { useProjectSettingStore } from "@/store/modules/projectSetting";
  import { List } from "@/api/printer";

  const rules = ref({});
  const group = ref('spaprintersetting');
  const show = ref(false);
  const formRef: any = ref(null);
  const message = useMessage();
  const formBtnLoading = ref(false);
  const settingStore = useProjectSettingStore();

  const printerList = ref([])
  const printerIdsArr = ref([])

  const formValue = ref({
    isOpen: 2,
    printerIds: '',
    printerNum: 1,
  });


  function formSubmit() {
    formBtnLoading.value = true;
    formRef.value.validate((errors) => {
      if (!errors) {
        formValue.value.printerIds = printerIdsArr.value ? printerIdsArr.value.join(',') : ''
        updateConfig({ group: group.value, list: formValue.value }).then((_res) => {
          formBtnLoading.value = false;
          message.success('更新成功');
          load();
        }).catch((err) => {
          formBtnLoading.value = false;
        });
      } else {
        message.error('验证失败，请填写完整信息');
        formBtnLoading.value = false;
      }
    });
  }

  onMounted(() => {
    load();
  });

  function load() {
    show.value = true;

    // 获取打印机列表
    List({
      pagination: false,
    }).then((res) => {
      printerList.value = res.list;
    });

    new Promise((_resolve, _reject) => {
      getConfig({ group: group.value })
        .then((res) => {
          if(res.list){
            formValue.value = res.list;
            printerIdsArr.value = res.list.printerIds ? res.list.printerIds.split(',').map(item => Number(item)) : [];
          }
        })
        .finally(() => {
          show.value = false;
        });
    });
  }
</script>
