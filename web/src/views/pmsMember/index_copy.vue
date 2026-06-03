<template>
  <n-loading-bar-provider>
    <div class="flex-column" style="height: calc(100vh - 118px)">
      <n-card :bordered="false" title="会员信息">
        <!--  这是由系统生成的CURD表格，你可以将此行注释改为表格的描述 -->
        <div class="flex-row">
          <n-input
            v-model:value="state.search.memberNo"
            type="text"
            placeholder="会员号"
            class="mr-2"
            style="width: 180px"
            @keyup.enter="loadDataTable"
          />
          <n-input
            v-model:value="state.search.fullName"
            type="text"
            placeholder="会员名"
            class="mr-2"
            style="width: 180px"
            @keyup.enter="loadDataTable"
          />
          <n-input
            v-model:value="state.search.phone"
            type="text"
            placeholder="手机号"
            class="mr-2"
            style="width: 180px"
            @keyup.enter="loadDataTable"
          />
          <n-button type="primary" class="ml-1" @click="loadDataTable"> 搜索 </n-button>
        </div>
      </n-card>

      <div class="flex-item pl-5 bgfff">
        <a-table
          :pagination="false"
          :dataSource="state.dataSource"
          :columns="state.columns"
          bordered
          size="middle"
        >
          <template #bodyCell="{ column, record }">
            <template v-if="column.key === 'memberNo'">
              <div class="flex-row">
                <img
                  v-if="record.avatar"
                  :src="record.avatar"
                  style="width: 45px; height: 45px; border-radius: 50%"
                />
                <img
                  v-else
                  src="~@/assets/images/mrtx.png"
                  style="width: 45px; height: 45px; border-radius: 50%"
                />
                <div class="flex-item" style="padding-left: 10px; line-height: 22px">
                  <div>
                    <span class="cblue f14 fw">{{ record.memberNo }}</span>
                    <a class="c999 f12 ml-1" v-copy="record.memberNo">复制</a>
                    <span class="c999 ml-2"
                      >level：<span class="cwarning">{{ record.level }}</span></span
                    >
                  </div>
                  <div>{{ record.fullName }}</div>
                </div>
              </div>
            </template>

            <template v-if="column.key === 'source'">
              <div class="lh-2">
                <n-icon size="22" color="#9E9E9E" v-if="record.source == 'IOS'">
                  <LogoApple />
                </n-icon>
                <n-icon size="22" color="#9E9E9E" v-else-if="record.source == 'H5'">
                  <LogoAndroid />
                </n-icon>
                <n-icon size="22" color="#2c7b05" v-else>
                  <LogoAndroid />
                </n-icon>
              
              </div>
            </template>
            <template v-if="column.key === 'phone'">
              <div>
                <span class="c999 mr-1">tel:</span>
                <span v-if="record.phone">{{ record.phoneArea }} - {{ record.phone }}</span>

                <a v-if="record.phone" class="c999 f12 ml-1" v-copy="record.phone">复制 </a>
              </div>
              <div>
                <span class="c999 mr-1">mail:</span>{{ record.mail }}
                <a v-if="record.mail" class="c999 f12 ml-1" v-copy="record.mail">复制</a></div
              >
            </template>

            <template v-if="column.key === 'work'">
              <div>
                <n-button type="primary" class="mr-2" size="small" @click="handleEdit(record)">
                  编辑
                </n-button>
                <n-button type="primary" class="mr-2" size="small" @click="handleView(record)">
                  详情
                </n-button>
                <n-dropdown trigger="click" :options="state.options" @select="handleSelect">
                  <n-button @click="selectclick(record)">通知</n-button>
                </n-dropdown>
              </div>
            </template>
          </template>
        </a-table>
      </div>
      <div class="text-r p-10">
        <a-pagination
          v-model:current="state.search.page"
          :total="state.totalCount"
          simple
          @change="changepage"
        />
      </div>

      <Edit ref="editRef" @reloadTable="reloadTable" />
      <View ref="viewRef" />
      <sendemail ref="sendemailref" />
      <sendnotify ref="sendnotifyref" />
      <sendsms ref="sendsmsref" />
    </div>
  </n-loading-bar-provider>
</template>

<script lang="ts" setup>
  import { h, reactive, ref, onMounted } from 'vue';
  import { useDialog, useMessage } from 'naive-ui';

  import { usePermission } from '@/hooks/web/usePermission';
  import { List, Export } from '@/api/pmsMember';

  import Edit from './edit.vue';
  import View from './view.vue';
  import sendemail from '@/views/smjcomm/sendemail.vue';
  import sendnotify from '@/views/smjcomm/sendnotify.vue';
  import sendsms from '@/views/smjcomm/sendsms.vue';

  import { Dicts } from '@/api/dict/dict';
  import { useLoadingBar } from 'naive-ui';
  import { LogoApple, LogoAndroid, GlobeSharp } from '@vicons/ionicons5';

  const dialog = useDialog();
  const message = useMessage();
  const actionRef = ref();
  const editRef = ref();
  const viewRef = ref();
  const sendemailref = ref();
  const sendnotifyref = ref();
  const sendsmsref = ref();

  const showreturn = ref(false);
  const loadingBar = useLoadingBar();

  const state = reactive({
    rowobj: null,
    btnloading: false,
    options: [
      {
        label: '发通知',
        key: 'mes',
      },
      {
        label: '发邮件',
        key: 'mail',
      },
      {
        label: '发短信',
        key: 'telmes',
      },
    ],
    totalCount: 0,
    search: {
      memberNo: '',
      fullName: '',
      phone: '',
      page: 1,
      pageSize: 10,
    },
    dataSource: [],
    columns: [
      {
        title: '会员信息',
        dataIndex: 'memberNo',
        key: 'memberNo',
      },
      {
        title: '手机/邮箱',
        dataIndex: 'phone',
        key: 'phone',
      },
      {
        title: '经验',
        dataIndex: 'exp',
        key: 'exp',
      },
      {
        title: '积分',
        dataIndex: 'balance',
        key: 'balance',
      },

      {
        title: '注册来源',
        dataIndex: 'source',
        key: 'source',
      },

      {
        title: '上次登录时间',
        dataIndex: 'lastLogin',
        key: 'lastLogin',
      },
      {
        title: '创建时间',
        dataIndex: 'createdAt',
        key: 'createdAt',
      },
      {
        title: '操作',
        dataIndex: 'work',
        key: 'work',
      },
    ],
  });
  const selectclick = (row) => {
    state.rowobj = row;
  };
  const handleSelect = (key: string | number) => {
    if (key === 'view') {
      return handleView(state.rowobj);
    } else if (key === 'mes') {
      return tomessage(state.rowobj);
    } else if (key === 'mail') {
      return goemail(state.rowobj);
    } else if (key === 'telmes') {
      return gosend(state.rowobj);
    }
  };
  function tomessage(aa) {
    sendnotifyref.value.openModal(aa.id);
  }
  function gosend(aa) {
    if (aa.phone) {
      sendsmsref.value.openModal({
        phone: aa.phone,
        area_no: aa.phoneArea,
      });
    } else {
      message.error('未找到联络方式');
    }
  }
  function goemail(aa) {
    if (aa.mail) {
      sendemailref.value.openModal(aa.mail);
    } else {
      message.error('改会员无邮箱');
    }
  }
  // 加载表格数据
  const loadDataTable = async () => {
    loadingBar.start();
    let data = JSON.parse(JSON.stringify(state.search));

    const res = await List(data);
    state.dataSource = res.list;
    state.totalCount = res.totalCount;
    loadingBar.finish();
  };

  // 重新加载表格数据
  function reloadTable() {
    actionRef.value?.reload();
  }

  // 编辑数据
  function handleEdit(record: Recordable) {
    editRef.value.openModal(record);
  }

  // 查看详情
  function handleView(record: Recordable) {
    viewRef.value.openModal(record);
  }
  function changepage(page, pageSize) {
    state.search.page = page;
    loadDataTable();
  }

  onMounted(() => {
    loadDataTable();
  });
</script>

<style lang="less" scoped></style>
