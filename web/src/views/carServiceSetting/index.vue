<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-card :bordered="false" :header-style="{
                    padding: '20px',
                  }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">其他基础配置</text>
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
              <div style="margin-bottom: 5px;">
                <template style="display:inline-block;">
                  <n-form-item label="增项服务" path="childrenSeatAddPrice">
                    <n-input-group>
                      <n-input-group-label>1个儿童座椅加收费用</n-input-group-label>
                      <n-input-number placeholder="请输入" :min="0" :show-button="false" v-model:value="formValue.childrenSeatAddPrice" style="width: 100px" />
                      <n-input-group-label>JPY</n-input-group-label>
                    </n-input-group>
                    <n-input-group style="margin-left: 8px">
                      <n-switch v-model:value="formValue.childrenSeatIsEnable" :unchecked-value="2" :checked-value="1">
                        <template #checked>
                          开启
                        </template>
                        <template #unchecked>
                          关闭
                        </template>
                      </n-switch>
                    </n-input-group>
                    <template #feedback>
                      <div>填0则免费;</div>
                      <div>开启后，用户可以在下单时选择使用儿童座椅；关闭则不显示</div>
                    </template>
                  </n-form-item>
                </template>
              </div>

              <div>
                <template style="display:inline-block;">
                  <n-form-item label=" " path="pickUpSignPrice">
                <n-input-group>
                  <n-input-group-label>举牌接机服务费用</n-input-group-label>
                  <n-input-number placeholder="请输入" :min="0" :show-button="false" v-model:value="formValue.pickUpSignPrice" style="width: 100px" />
                  <n-input-group-label>JPY</n-input-group-label>
                </n-input-group>
                <n-input-group style="margin-left: 8px">
                  <n-switch v-model:value="formValue.pickUpSignIsEnable" :unchecked-value="2" :checked-value="1">
                    <template #checked>
                      开启
                    </template>
                    <template #unchecked>
                      关闭
                    </template>
                  </n-switch>
                </n-input-group>
                <template #feedback>
                  <div>填0则免费;</div>
                  <div>开启后，用户可以在下单时选择使用接机服务；关闭则不显示</div>
                </template>
              </n-form-item>
                </template>
              </div>
            </n-gi>
            <n-gi span="1">
                <n-form-item label="限时免费" path="limitTimeFreeArr">
                  <n-checkbox-group v-model:value="formValue.limitTimeFreeArr">
                    <n-space>
                      <n-checkbox
                        :value="1"
                        label="1个儿童座椅加收费用"
                      />
                      <n-checkbox
                        :value="2"
                        label="举牌接机服务费用"
                      />
                    </n-space>
                  </n-checkbox-group>
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
import {ref, onMounted} from 'vue';
  import { useMessage } from 'naive-ui';
  import { getConfig, updateConfig } from '@/api/sys/config';
  import {useProjectSettingStore} from "@/store/modules/projectSetting";
import {Text} from "@/api/translate";

  const rules = ref({});
  const group = ref('carservicesetting');
  const show = ref(false);
  const formRef: any = ref(null);
  const message = useMessage();
  const formBtnLoading = ref(false);
  const settingStore = useProjectSettingStore();

  const formValue = ref({
    childrenSeatAddPrice: 0,
    pickUpSignPrice: 0,
    limitTimeFree: '',
    limitTimeFreeArr: [],
    childrenSeatIsEnable: 0,
    pickUpSignIsEnable: 0,
  });

  function formSubmit() {
    formBtnLoading.value = true;
    formRef.value.validate((errors) => {
      if (!errors) {
        formValue.value.limitTimeFree = formValue.value.limitTimeFreeArr.join(',')
        delete formValue.value.limitTimeFreeArr
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
    new Promise((_resolve, _reject) => {
      getConfig({ group: group.value })
        .then((res) => {
          if(res.list){
            formValue.value = res.list;
            formValue.value.limitTimeFreeArr = res.list.limitTimeFree ? res.list.limitTimeFree.split(',').map((item)=>{
              return parseInt(item)
            }) : [];
          }
          show.value = false;
        }).catch((err)=>{
          show.value = false;
        })
    });
  }
</script>
