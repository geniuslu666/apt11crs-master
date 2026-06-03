<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-card :bordered="false" :header-style="{
                    padding: '20px',
                  }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">基础设置</text>
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
              <n-form-item label="是否开启" path="isOpen">
                <n-switch v-model:value="formValue.isOpen" :unchecked-value="2" :checked-value="1"/>
              </n-form-item>
            </n-gi>

            <n-gi span="1" style="margin-bottom: 25px" v-if="formValue.isOpen == 1">
              <n-form-item label="允许调试类型" path="testType">
                <n-radio-group v-model:value="formValue.testType" name="cancelPolicyOpen">
                  <n-space>
                    <n-radio :value="1">
                      会员分组
                    </n-radio>
                    <n-radio :value="2">
                      指定会员
                    </n-radio>
                  </n-space>
                </n-radio-group>
                <template #feedback>指定分组/会员可以在开启维护模式后，在APP端访问应用</template>
              </n-form-item>
            </n-gi>

            <n-gi span="1" v-if="formValue.isOpen == 1&&formValue.testType == 1">
              <n-form-item label="会员分组" path="group" style="width: 800px">
                <n-select
                  placeholder="请选择会员分组"
                  v-model:value="canTestGroupIdsArr"
                  :options="groupList"
                  label-field="memberGroup"
                  value-field="id"
                  clearable
                  filterable
                  multiple
                  style="width: 300px"
                />
              </n-form-item>
            </n-gi>
            <n-gi span="1" v-if="formValue.isOpen == 1&&formValue.testType == 2">
              <n-form-item label="会员" path="group" style="width: 800px">
                <n-button type="primary" @click="chooseMemberBtn">选择会员</n-button>
              </n-form-item>
              <div style="margin-left: 40px;margin-top: 10px;margin-bottom: 20px">
                <n-table>
                  <thead>
                  <tr>
                    <th>会员号</th>
                    <th>会员全称</th>
                    <th>电话</th>
                    <th>操作</th>
                  </tr>
                  </thead>
                  <tbody>
                  <tr v-for="(item, index) of chooseMemberInfoArr" :key="index">
                    <td>{{ item.memberNo }}</td>
                    <td>
                      {{ item.fullName }}
                    </td>
                    <td>{{ item.phoneArea }}-{{ item.phone }}</td>
                    <td>
                      <n-button type="error" @click="handleDeleteCoupon(index)">删除</n-button>
                    </td>
                  </tr>
                  </tbody>
                </n-table>
              </div>
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
    <ChooseMember ref="chooseMemberRef" @reloadMemberList="handleChooseMemberInfo"/>
  </div>
</template>

<script lang="ts" setup>
  import { ref, onMounted } from 'vue';
  import { useMessage } from 'naive-ui';
  import { getConfig, updateConfig } from '@/api/sys/config';
  import {useProjectSettingStore} from "@/store/modules/projectSetting";
  import ChooseMember from "@/components/ChooseMember/chooseMember.vue";
  import {List} from "@/api/pmsMember";
  import {GroupAll} from "@/api/pmsMemberGroup";

  const rules = ref({});
  const group = ref('foodmaintenancebasic');
  const show = ref(false);
  const formRef: any = ref(null);
  const message = useMessage();
  const formBtnLoading = ref(false);
  const settingStore = useProjectSettingStore();

  const canTestGroupIdsArr = ref([]);
  const canTestMemberIdsArr = ref([]);
  const groupList = ref([])
  const chooseMemberRef = ref();
  const chooseMemberInfoArr = ref([]);

  const formValue = ref({
    isOpen: 2,
    testType: 1,
    canTestGroupIds: '',
    canTestMemberIds: '',
  });

  function handleChooseMemberInfo(records){

    console.log('records------------',records)

    chooseMemberInfoArr.value.push(records)
    canTestMemberIdsArr.value.push(records.id)
    formValue.value.canTestMemberIds = canTestMemberIdsArr.value.join(',')

  }

  function chooseMemberBtn(){
    chooseMemberRef.value.openModal(canTestMemberIdsArr.value);
  }

  function handleDeleteCoupon(index){
    chooseMemberInfoArr.value.splice(index, 1);
    canTestMemberIdsArr.value.splice(index, 1);
    formValue.value.canTestMemberIds = canTestMemberIdsArr.value.join(',')
  }

  // 获取会员列表
  function getMemberList(){
    List({
      memberIds: formValue.value.canTestMemberIds,
      pagination: false
    }).then((res) => {
      chooseMemberInfoArr.value = res.list;
    })
      .finally(() => {
        show.value = false;
      });
  }

  function formSubmit() {
    formBtnLoading.value = true;
    formRef.value.validate((errors) => {
      if (!errors) {

        if(formValue.value.testType == 1){
          formValue.value.canTestGroupIds = canTestGroupIdsArr.value ? canTestGroupIdsArr.value.join(',') : ''
          formValue.value.canTestMemberIds = ''
        }
        if(formValue.value.testType == 2){
          formValue.value.canTestGroupIds = ''
        }

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
    GroupAll({}).then((res) => {
      groupList.value = res.list;
    });
    new Promise((_resolve, _reject) => {
      getConfig({ group: group.value })
        .then((res) => {
          if(res.list){
            formValue.value = res.list;
            canTestGroupIdsArr.value = res.list.canTestGroupIds ? res.list.canTestGroupIds.split(',').map(item => Number(item)) : [];
            canTestMemberIdsArr.value = res.list.canTestMemberIds ? res.list.canTestMemberIds.split(',') : [];
          }
          if(res.list.canTestMemberIds != ''){
            getMemberList();
          }else{
            chooseMemberInfoArr.value = [];
          }
        }).finally(() => {
          show.value = false;
        });
    });
  }
</script>
