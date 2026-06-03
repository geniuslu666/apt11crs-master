<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-form
        ref="formRef"
        :model="formValue"
        :rules="rules"
        :label-placement="settingStore.isMobile ? 'top' : 'left'"
        :label-width="150"
        class="py-4"
      >
        <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
          <n-gi span="1">
            <n-form-item label="司机昵称" path="nickname">
              <n-input placeholder="请输入司机昵称" v-model:value="formValue.nickname" style="width: 150px" />
            </n-form-item>
          </n-gi>
          <n-gi span="1">
            <n-form-item label="真实姓名" path="name">
              <n-input placeholder="请输入真实姓名" v-model:value="formValue.name" style="width: 150px" />
            </n-form-item>
          </n-gi>
          <n-gi span="1">
            <n-form-item label="性别" path="sex" :show-require-mark="true">
              <n-radio-group v-model:value="formValue.sex">
                <n-radio-button
                  :value="1"
                  label="男"
                />
                <n-radio-button
                  :value="2"
                  label="女"
                />
              </n-radio-group>
            </n-form-item>
          </n-gi>
          <n-gi span="1">
            <n-form-item label="手机号码" path="phone" :show-require-mark="true">
              <n-input-group>
                <n-input placeholder="区号" v-model:value="formValue.phoneArea" style="width: 50px" />
                <n-input-group-label>-</n-input-group-label>
                <n-input placeholder="手机号" v-model:value="formValue.phone" style="width: 150px" />
              </n-input-group>
            </n-form-item>
          </n-gi>
          <n-gi span="1">
            <n-form-item label="照片" path="photo" :show-require-mark="true" :show-feedback='false'>
              <UploadImage :maxNumber="1" v-model:value="formValue.photo" />
            </n-form-item>
            <div style="color:#9EA4AA; font-size: 12px;line-height: 34px;margin: 0 0 24px;padding-left: 150px;">建议尺寸：128px*174px</div>
          </n-gi>
          <n-gi span="1">
            <n-form-item label="年龄" path="age" :show-require-mark="true">
              <n-input-group>
                <n-input-number placeholder="请输入" :show-button="false" :min="0" :precision="0" v-model:value="formValue.age" style="width: 100px" />
                <n-input-group-label>岁</n-input-group-label>
              </n-input-group>
            </n-form-item>
          </n-gi>
          <n-gi span="1">
            <n-form-item label="从业年数" path="workYears" :show-require-mark="true">
              <n-input-group>
                <n-input-number placeholder="请输入" :show-button="false" :min="0" :precision="0" v-model:value="formValue.workYears" style="width: 100px" />
                <n-input-group-label>年</n-input-group-label>
              </n-input-group>
            </n-form-item>
          </n-gi>
          <n-gi span="1">
            <n-form-item label="语言能力" path="language" :show-require-mark="true">
              <n-checkbox-group v-model:value="formValue.languageArr">
                <n-space>
                  <n-checkbox
                    v-for="item in languageArr"
                    :value="item.value"
                    :label="item.label"
                  />
                </n-space>
              </n-checkbox-group>
            </n-form-item>
          </n-gi>
          <n-gi span="1">
            <n-form-item label="合作类型" path="cooperateTypeId" :show-require-mark="true">
              <n-radio-group v-model:value="formValue.cooperateTypeId">
                <n-radio-button
                  v-for="item in cooperateTypeList"
                  :key="item.id"
                  :value="item.id"
                  :label="item.typeName"
                />
              </n-radio-group>
            </n-form-item>
          </n-gi>
          <n-gi span="1">
            <n-form-item label="状态" path="status" :show-require-mark="true">
              <n-radio-group v-model:value="formValue.status" name="status">
                <n-radio-button
                  v-for="type in options.sys_normal_disable"
                  :key="type.value"
                  :value="type.value"
                  :label="type.label"
                />
              </n-radio-group>
            </n-form-item>
          </n-gi>
        </n-grid>
        <div style="text-align: center;margin-top: 30px">
          <n-space justify="center">
            <n-button type="info" :loading="formBtnLoading" @click="confirmForm">
              确定
            </n-button>
          </n-space>
        </div>
      </n-form>
    </n-spin>
  </div>
</template>
<script setup lang="ts">
import {onMounted, ref, watch} from "vue";
import {useProjectSettingStore} from "@/store/modules/projectSetting";
import UploadImage from "@/components/Upload/uploadImage.vue";
import {Edit} from "@/api/carDriver";
import {useMessage} from "naive-ui";
import {useTabsViewStore} from "@/store/modules/tabsView";
import {useRouter} from "vue-router";
import {loadOptions, options} from "@/views/carDriver/model";
import {List as CooperateTypeList} from "@/api/carCooperateType";

const props = defineProps({
  formData: {
    type: Object || null,
    default: null,
  },
});
const emit = defineEmits(['reloadInfo']);
const router = useRouter();
const tabsViewStore = useTabsViewStore();
const show = ref(false);
const formValue = ref({
  type: 'basic',
  id: 0,
  nickname: '',
  sex: 1,
  name: '',
  phoneArea: '',
  phone: '',
  photo: '',
  age: null,
  workYears: null,
  status: 2,
  cooperateTypeId: null,
  language: '',
  languageArr: []
});
const settingStore = useProjectSettingStore();
const formBtnLoading = ref(false);
const formRef = ref<any>({});
const message = useMessage();
const languageArr = ref([
  {
    label: '中文',
    value: 'zh'
  },
  {
    label: '英文',
    value: 'en'
  },
  {
    label: '日语',
    value: 'ja'
  },
  {
    label: '韩语',
    value: 'ko'
  }
])
const cooperateTypeList = ref([]);

const rules = ref({
  nickname: {
    required: true,
    trigger: ['blur', 'input'],
    message: '请输入昵称'
  },
  name: {
    required: true,
    trigger: ['blur', 'input'],
    message: '请输入真实姓名'
  },
});

function confirmForm(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      if(!formValue.value.phoneArea || !formValue.value.phone){
        formBtnLoading.value = false;
        message.error('手机号请填写完整');
        return false;
      }
      if(!formValue.value.photo){
        formBtnLoading.value = false;
        message.error('请上传照片');
        return false;
      }
      if(!formValue.value.age){
        formBtnLoading.value = false;
        message.error('请输入年龄');
        return false;
      }
      if(!formValue.value.workYears){
        formBtnLoading.value = false;
        message.error('请输入从业年数');
        return false;
      }
      if(formValue.value.languageArr.length <= 0){
        formBtnLoading.value = false;
        message.error('请选择语言能力');
        return false;
      }
      if(!formValue.value.cooperateTypeId){
        formBtnLoading.value = false;
        message.error('请选择合作类型');
        return false;
      }
      formValue.value.language = formValue.value.languageArr.join(',')

      Edit(formValue.value).then((_res) => {
        message.success('操作成功');
        formBtnLoading.value = false;
        setTimeout(() => {
          if(formValue.value.id > 0){
            emit('reloadInfo');
          }else{
            // todo 关闭当前tab页，并跳转到等级列表页
            setTimeout(() => {
              tabsViewStore.closeSignal('2');
              router.push({ name: 'carDriverIndex', params: {  } });
            }, 500);

          }
        }, 500);
      }).catch((err) => {
        formBtnLoading.value = false;
      });
    } else {
      formBtnLoading.value = false;
      message.error('请填写完整信息');
    }
  });
}

async function loadCooperateTypeList(){
  let cooperateTypeArrListOrg = await CooperateTypeList({
    status: 1,
    Pagination: false
  })
  cooperateTypeList.value = cooperateTypeArrListOrg.list
}

async function initData(data){
  formValue.value = {
    type: 'basic',
    id: data.id,
    nickname: data.nickname,
    name: data.name,
    sex: data.sex,
    phoneArea: data.phoneArea,
    phone: data.phone,
    photo: data.photo,
    age: data.age,
    workYears: data.workYears,
    status: data.status,
    cooperateTypeId: data.cooperateTypeId,
    language: data.language,
    languageArr: data.language ? data.language.split(',') : [],
  }
}

watch(
  () => props.formData,
  async(value) => {
    await initData(value)
  }
);

onMounted(async() => {
  show.value = true;
  await loadCooperateTypeList()
  await initData(props.formData)
  await loadOptions();
  show.value = false;
});

</script>

<style lang="less">
.mapdiv {
  position: absolute;
  top: 15px;
  left: 15px;
  height: 38px;
  text-align: center;
  color: #3d3d3d;
  padding: 8px 15px;
  background-color: #fff;
  box-shadow: 4px 3px 5px #999999a8;
  z-index: 300;
  .tabview {
    color: #333;
    font-size: 14px;
    .showmap {
      color: #053dc8;
      font-weight: 550;
    }
    img {
      width: 18px;
      height: 18px;
      float: left;
      margin-right: 6px;
    }
  }
}
.mapbtm {
  position: absolute;
  z-index: 300;
  bottom: 0;
  text-align: left;
  width: 100%;
  background: #6a7b8c38;
  padding: 5px 10px;
  font-weight: 550;
}
</style>
