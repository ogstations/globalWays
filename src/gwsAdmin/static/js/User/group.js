
/**
 * 会员组管理
 */


//添加会员组
function add() {
	art.dialog.open('/Group/add/', {
		id : 'group_add',
		title : '添加用户组',
		width : 880,
		height : 300,
		lock : true,
		ok : function() {
			var iframe = this.iframe.contentWindow;

			var par = [];

			var name = iframe.$('#name').val();
			if (name == '') {
				iframe.$('#nametip').removeClass("onShow").addClass("onError").html('请输入会员组名称!');
				return false;
			}else{
				iframe.$('#nametip').removeClass("onError").addClass("onCorrect").html('输入正确');
				var pars = "name=" + name;
				par.push(pars);
			}

			var desc = iframe.$('#desc').val();
			if (desc == '') {
				iframe.$('#desctip').removeClass("onShow").addClass("onError").html('请输入会员组描述!');
				return false;
			}else{
				iframe.$('#desctip').removeClass("onError").addClass("onCorrect").html('输入正确');
				var pars = "desc=" + desc;
				par.push(pars);
			}

			var contribution = iframe.$('#contribution').val();
			if (contribution == '') {
				iframe.$('#contributiontip').removeClass("onShow").addClass("onError").html('请输入会费!');
				return false;
			}else{
				iframe.$('#contributiontip').removeClass("onError").addClass("onCorrect").html('输入正确');
				var pars = "contribution=" + contribution;
				par.push(pars);
			}

			var status = iframe.$("input[name='status']:checked").val();
			var pars = "status=" + status;
			par.push(pars);

			pars = par.join("&");

			$.ajax({
				type : "POST",
				url : "/Group/add/",
				data : pars,
				success : function(tmp) {
					if (tmp.status == 1) {
						right_refresh();
						notice_tips("添加成功!");
					} else {
						notice_tips(tmp.message);
					}
				}
			});
		},
		cancel : true
	});
}

/**
 * 编辑用户组
 */
function edit(id) {
	if (id == '') {
		notice_tips("参数错误!");
		return false;
	}

	art.dialog.open('/Group/edit/' + id + '/', {
		id : 'group_edit',
		title : '编辑用户',
		width : 880,
		height : 300,
		lock : true,
		ok : function() {
			var iframe = this.iframe.contentWindow;

			var par = [];

			var pars = "id=" + id;
			par.push(pars);

			var name = iframe.$('#name').val();
            if (name == '') {
                iframe.$('#nametip').removeClass("onShow").addClass("onError").html('请输入会员组名称!');
                return false;
            }else{
                iframe.$('#nametip').removeClass("onError").addClass("onCorrect").html('输入正确');
                var pars = "name=" + name;
                par.push(pars);
            }

            var desc = iframe.$('#desc').val();
            if (desc == '') {
                iframe.$('#desctip').removeClass("onShow").addClass("onError").html('请输入会员组描述!');
                return false;
            }else{
                iframe.$('#desctip').removeClass("onError").addClass("onCorrect").html('输入正确');
                var pars = "desc=" + desc;
                par.push(pars);
            }

            var contribution = iframe.$('#contribution').val();
            if (contribution == '') {
                iframe.$('#contributiontip').removeClass("onShow").addClass("onError").html('请输入会费!');
                return false;
            }else{
                iframe.$('#contributiontip').removeClass("onError").addClass("onCorrect").html('输入正确');
                var pars = "contribution=" + contribution;
                par.push(pars);
            }

            var status = iframe.$("input[name='status']:checked").val();
            var pars = "status=" + status;
            par.push(pars);

			pars = par.join("&");

notice_tips("编辑用户组成功!");
			$.ajax({
				type : "POST",
				url : "/Group/edit/",
				data : pars,
				success : function(tmp) {
					if (tmp.status == 1) {
						right_refresh();
						notice_tips("编辑成功!");
					} else {
						notice_tips(tmp.message);
					}
				}
			});
		},
		cancel : true
	});
}

/**
 * 删除用户组
 */
function del(id) {
	if (id == '') {
		notice_tips("参数错误!");
		return false;
	}

	art.dialog.confirm('你确定要删除这个用户吗?', function() {
		$.ajax({
			type : "POST",
			url : "/Group/delete/",
			data : "id=" + id,
			success : function(tmp) {
				if (tmp.status == 1) {
					right_refresh();
					notice_tips("删除成功!");
				} else {
					notice_tips(tmp.content);
				}
			}
		});
	}, function() {
		notice_tips("你取消了删除用户操作!");
	});
}

function set_username_color(color) {
	$('#usernamecolor').val(color);
}