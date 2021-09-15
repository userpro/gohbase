###
# @Author: dongzhzheng
# @Date: 2021-09-15 20:13:44
# @LastEditTime: 2021-09-15 20:28:15
# @LastEditors: dongzhzheng
# @FilePath: /userpro-gohbase/pb/generate.sh
# @Description:
###

PROTO_FILE=""
for dir in $(find . -name "*.proto"); do PROTO_FILE="${PROTO_FILE} --go_opt M${dir:2}=../pb"; done
echo "generate proto: ${PROTO_FILE}"

GENERATE_FILE="generate.go"
echo "GENERATE_FILE: ${GENERATE_FILE}"
echo "package pb" >${GENERATE_FILE}
echo "//go:generate sh -c \"protoc ${PROTO_FILE} --go_out=. *.proto\"" >>${GENERATE_FILE}
