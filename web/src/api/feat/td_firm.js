import request from "@/utils/request";

const BASE_API = "/admin/feat";

export default {
  all(query) {
    return request({
      url: BASE_API + '/tdFirm/all',
      method: "get",
      params: query,
    });
  },
  page(query) {
    return request({
      url: BASE_API + '/tdFirm/page',
      method: "get",
      params: query,
    });
  },
  create(data) {
    return request({
      url: BASE_API + '/tdFirm/create',
      method: "post",
      data,
    });
  },
  update(data) {
    return request({
      url: BASE_API + '/tdFirm/update',
      method: "post",
      data,
    });
  },
  detail(id) {
      return request({
        url: BASE_API + '/tdFirm/detail',
        method: "get",
        params: { id: id },
      });
  },
  delete(id) {
    return request({
      url: BASE_API + '/tdFirm/delete',
      method: "post",
      data: { id: id },
    });
  },
  deletes(ids) {
    return request({
      url: BASE_API + '/tdFirm/deletes',
      method: "post",
      data: { ids: ids },
    });
  },
};