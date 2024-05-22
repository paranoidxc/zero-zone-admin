<template>
  <div style="margin: 20px">
    <div class="">
      <el-form :model="tableSearchForm" inline>
        <el-form-item label="厂商名称">
            <el-input v-model="tableSearchForm.firmName" placeholder="" clearable />
        </el-form-item>
        <el-form-item label="厂商别名">
            <el-input v-model="tableSearchForm.firmAlias" placeholder="" clearable />
        </el-form-item>
        <el-form-item label="厂商编码">
            <el-input v-model="tableSearchForm.firmCode" placeholder="" clearable />
        </el-form-item>
        <el-form-item label="厂商描述">
            <el-input v-model="tableSearchForm.firmDesc" placeholder="" clearable />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSearchSubmit">
            <el-icon class="el-icon--left">
              <Search />
            </el-icon>
            查询
          </el-button>
        </el-form-item>
      </el-form>
    </div>

    <!-- query -->
    <div class="query-box">
      <div class="btn-list">
        <el-button type="primary" @click="handleCreate">
          <el-icon class="el-icon--left">
            <Plus />
          </el-icon>
          增加
        </el-button>
        <el-button
          type="danger"
          @click="handleDelList"
          v-if="multipleSelection.length > 0"
        >
          <el-icon class="el-icon--left">
            <Delete />
          </el-icon>
          删除多选
        </el-button>
      </div>
      <div class="btn-list"></div>
    </div>

    <el-table
      ref="multipleTableRef"
      @selection-change="handleSelectionChange"
      :data="tableData"
      height="460"
      stripe
      style="width: 100%"
      border
    >
      <el-table-column fixed type="selection" width="55" />
      <el-table-column fixed prop="id" label="编号" width="60" />
      <el-table-column fixed prop="name" label="名称" width="100" />
      <el-table-column prop="orderNum" label="排序值" width="100" />
      <el-table-column prop="status" label="状态" width="60">
        <template #default="scope">
        </template>
      </el-table-column>
      <el-table-column prop="remark" label="备注" width="" />
      <el-table-column fixed="right" label="操作" width="160">
        <template #default="scope">
          <el-button size="small" @click="handleEdit(scope.row)">
            编 辑
          </el-button>
          <el-button
            size="small"
            type="danger"
            @click="handleRowDel(scope.row)"
          >
            删 除
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
      background
      style="display: flex; justify-content: center; margin-top: 10px"
      v-model:current-page="curPage"
      v-model:page-size="limit"
      :page-sizes="[limit, 10, 20, 50, 100, 200, 300, 400, 500]"
      layout="total, sizes, prev, pager, next, jumper"
      :total="total"
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
    />

    <!-- dialog -->
    <el-dialog
      draggable
      v-model="dialogFormVisible"
      :title="dialogType === 'add' ? '新增' : '编辑'"
    >
      <el-form :model="tableForm">
        <el-form-item
          style="display: none"
          v-if="dialogType === 'edit'"
          label="编号"
          :label-width="80"
        >
          <el-input v-model="tableForm.id" autocomplete="off" />
        </el-form-item>
        <el-form-item label="名称" :label-width="80">
          <el-input v-model="tableForm.name" autocomplete="off" />
        </el-form-item>

        <el-form-item label="状态" :label-width="80">
          <el-tooltip
            :content="getStatusLabel(tableForm.status)"
            placement="top"
          >
            <el-switch
              v-model.number="tableForm.status"
              class="mt-2"
              inline-prompt
              :active-value="1"
              :inactive-value="0"
            />
          </el-tooltip>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button type="primary" @click="dialogConfirm"> 确 认 </el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>
<script setup>
import { getCurrentInstance, proxyRefs } from "vue";

const { proxy } = getCurrentInstance();

import sysTableApi from "@/api/feat/td_firm.js";

let tableSearchForm = $ref({});
let tableData = $ref([]); // 表格数据
let tableForm = $ref({
  status: 1,
});
let dialogFormVisible = $ref(false);
let dialogType = $ref("add");
let multipleSelection = $ref([]);
let limit = $ref(2);
let total = $ref(15);
let curPage = $ref(1);

const statusOptions = [
  {
    value: 1,
    label: "启用",
  },
  {
    value: 0,
    label: "禁用",
  },
];

const getStatusLabel = (idx) => {
  const index = statusOptions.findIndex((option) => option.value === idx);
  if (index !== -1) {
    return statusOptions[index].label;
  } else {
  }
};

/* 请求分页 */
const getTableDataList = async (cur = 1, limit = 2) => {
  let result = await sysTableApi.page({ page: cur, limit: limit });
  if (result.code == 200) {
    tableData = result.data.list;
    curPage = result.data.page;
    total = result.data.total;
  }
};
getTableDataList();

const handleSizeChange = (val) => {
  limit = val;
  getTableDataList(curPage, val);
};

const handleCurrentChange = (val) => {
  getTableDataList(val, limit);
};

// 删除一条
const reqRowDel = async (id) => {
  await sysTableApi.delete(id);
};

const handleRowDel = async (id) => {
  await sysTableApi.delete(id);
  await getTableDataList(curPage);
};

const handleDelList = async() => {
  multipleSelection.forEach((id) => {
	reqRowDel(id)
  });
  multipleSelection = [];
  await getTableDataList(curPage);
};

// 选中
const handleSelectionChange = (val) => {
  multipleSelection = [];
  val.forEach((item) => {
    multipleSelection.push(item.id);
  });
};

// 编辑
const handleEdit = (row) => {
  dialogFormVisible = true;
  dialogType = "edit";
  tableForm = { ...row };
};

// 新增
const handleCreate = () => {
  dialogFormVisible = true;
  tableForm = {
    status: 1,
  };
  dialogType = "create";
};

//查询
const onSearchSubmit = async () => {
  tableSearchForm.page = 1;
  tableSearchForm.limit = limit;

  sysTableApi.listPage(tableSearchForm).then((res) => {
    if (res.code === 200) {
      tableData = res.data.list;
      curPage = res.data.page;
      total = res.data.total;
    }
  });
};

// 确认
const dialogConfirm = async () => {
  if (dialogType === "create") {
    // 添加数据
    tableForm.parentId = 0;
    tableForm.orderNum = 100;
    sysTableApi
      .create(tableForm)
      .then((res) => {
        if (res.code == 200) {
          dialogFormVisible = false;
          getTableDataList(curPage);
        }
      })
      .catch(() => {
        // console.log("CCCC");
      });
  } else {
    // 修改 内容
    sysTableApi.update(tableForm).then((res) => {
      if (res.code == 200) {
        dialogFormVisible = false;
        getTableDataList(curPage);
      }
    });
  }
};
</script>

<style scoped>
.query-box {
  display: flex;
  justify-content: space-between;
  margin-bottom: 10px;
}
</style>