<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" :header-style="{
                    padding: '20px',
                  }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">{{ translang('取消政策') }}</text>
        </template>
        <!--  这是由系统生成的CURD表格，你可以将此行注释改为表格的描述 -->
      </n-card>
    </div>
    <n-card :bordered="false" class="proCard" :content-style="{
                    padding: '0 20px 44px',
                  }">
      <n-table>
        <thead>
        <tr>
          <th>类型</th>
          <th>名称</th>
          <th>时间</th>
          <th>取消费</th>
        </tr>
        </thead>
        <tbody>
        <tr>
          <td>早于</td>
          <td>
            <div style="display: flex;align-items: center">
              <span v-if="state.before.nameLanguage.zh">{{ state.before.nameLanguage.zh.content }}</span>
              <n-button text tag="a" type="primary" style="margin-left: 8px;font-size: 12px"
                        @click="editItem(state.before.id,'CancelRateeditName','编辑取消政策名称',state.before.nameLanguage)">
                <img src="@/assets/images/icon_translate.png" width="16" style="margin-right: 3px">
                翻译
              </n-button>
            </div>
          </td>
          <td>
            <a-input-number
              id="inputNumber"
              v-model:value="state.before.endDays"
              addon-after="天"
              :max="500"
              :min="1"
              style="width: 120px"
              @change="beforenumberchange"
            />
          </td>
          <td>
            <a-input-number
              v-model:value="state.before.rate"
              addon-after="%"
              style="width: 200px"
              @change="pmsCancelRateeditchange(state.before)"
            ></a-input-number>
          </td>
        </tr>
        <tr>
          <td>{{ middleName }}</td>
          <td>
            <div style="display: flex;align-items: center">
              <span v-if="state.middle.nameLanguage.zh">{{ state.middle.nameLanguage.zh.content }}</span>
              <n-button text tag="a" type="primary" style="margin-left: 8px;font-size: 12px"
                        @click="editItem(state.middle.id,'CancelRateeditName','编辑取消政策名称',state.middle.nameLanguage)">
                <img src="@/assets/images/icon_translate.png" width="16" style="margin-right: 3px">
                翻译
              </n-button>
            </div>
          </td>
          <td>
            <a-input-number
              id="inputNumber"
              v-model:value="state.middle.endDays"
              addon-after="天"
              :max="500"
              :min="1"
              style="width: 120px"
              @change="middlenumberchange"
            />
          </td>
          <td>
            <a-input-number
              v-model:value="state.middle.rate"
              addon-after="%"
              style="width: 200px"
              @change="pmsCancelRateeditchange(state.middle)"
            ></a-input-number>
          </td>
        </tr>
        <tr>
          <td>{{ afterName }}</td>
          <td>
            <div style="display: flex;align-items: center">
              <span v-if="state.after.nameLanguage.zh">{{ state.after.nameLanguage.zh.content }}</span>
              <n-button text tag="a" type="primary" style="margin-left: 8px;font-size: 12px"
                        @click="editItem(state.after.id,'CancelRateeditName','编辑取消政策名称',state.after.nameLanguage)">
                <img src="@/assets/images/icon_translate.png" width="16" style="margin-right: 3px">
                翻译
              </n-button>
            </div>
          </td>
          <td></td>
          <td>
            <a-input-number
              v-model:value="state.after.rate"
              addon-after="%"
              style="width: 200px"
              @change="pmsCancelRateeditchange(state.after)"
            ></a-input-number>
          </td>
        </tr>
        </tbody>
      </n-table>
<!--      <div style="text-align: center;margin-top: 60px">-->
<!--        <n-space justify="center">-->
<!--          <n-button type="primary" :loading="formBtnLoading" @click="formSubmit">保存更新</n-button>-->
<!--        </n-space>-->
<!--      </div>-->
    </n-card>
    <baseEdit ref="editBaseRef" @reloadTable="DataLoad"/>
  </div>
</template>

<script lang="ts" setup>
import {onMounted, reactive, ref} from 'vue';
import baseEdit from '@/views/pmsProperty/comm/edit_base.vue';
import {useMessage} from 'naive-ui';
import {jsontoobj, translang} from '@/utils/smjcomm';
import {pmsCancelRateedit, pmsCancelRatelist} from '@/api/comm';

const message = useMessage();
const editBaseRef = ref();
const middleName = ref('从 0 天到');
const afterName = ref('晚于 0 天');
const state = reactive({
  List: [],
  before: {
    id: 0,
    date: 1,
    mode: 'string',
    startDays: 0,
    endDays: 0,
    rate: 0,
    name: '早于',
    nameLanguage: {}
  },
  middle: {
    id: 0,
    mode: '',
    startDays: 0,
    endDays: 0,
    rate: 0,
    name: '请选择',
    nameLanguage: {}
  },

  after: {
    id: 0,
    mode: '',
    startDays: 0,
    rate: 0,
    name: '请选择',
    nameLanguage: {}
  },
});
const DataLoad = async () => {
  const res = await pmsCancelRatelist();
  state.List = res.list;
  res.list.forEach((f) => {

    if (f.nameLanguage) {
      f.nameLanguage = jsontoobj(f.nameLanguage);
      if (f.nameLanguage.zh_CN) {
        f.nameLanguage.zh_CN = f.nameLanguage.zh_CN
      } else if (f.nameLanguage.zh_cn) {
        f.nameLanguage.zh_CN = f.nameLanguage.zh_cn
      } else {
        f.nameLanguage.zh_CN = {}
      }
    } else {
      f.nameLanguage = {};
    }

    console.log(f)
    if (f.mode == 'before') {
      state.before = f;
      middleName.value = '从 ' + (state.before.endDays - 1) + ' 天到'
    } else if (f.mode == 'middle') {
      state.middle = f;
      afterName.value = '晚于 ' + (state.middle.endDays - 1) + ' 天';
    } else if (f.mode == 'after') {
      state.after = f;
    }


  });
};
const beforenumberchange = (val) => {
  state.middle.startDays = val - 1;
  middleName.value = '从 ' + state.middle.startDays + ' 天到';
  pmsCancelRateeditchange(state.before);
  pmsCancelRateeditchange(state.middle);
};
const middlenumberchange = (val) => {
  state.after.startDays = val - 1;
  afterName.value = '晚于 ' + state.after.startDays + ' 天';
  pmsCancelRateeditchange(state.middle);
  pmsCancelRateeditchange(state.after);
};
const pmsCancelRateeditchange = async (data) => {

  // debugger
  // 创建一个新的对象，以避免直接修改原始数据
  var pmsCancelRateeditchangeparams = { ...data };
  delete pmsCancelRateeditchangeparams.nameLanguage;
  const res = await pmsCancelRateedit(pmsCancelRateeditchangeparams);
  DataLoad();
};
onMounted(() => {
  DataLoad();
});

function editItem(id, type, title, data) {
  editBaseRef.value.openModal(id, type, title, data);
}

</script>

<style lang="less">
.center_box {
  display: flex;
  align-items: center;
}
</style>
