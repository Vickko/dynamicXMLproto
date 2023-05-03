# dynamicXMLproto
fully-automatic xml2proto convertor using reflect based on go.


### TODOS

1. refactor logger using middleware design
2. protoDescFile Design
3. proto desc go impl


### TIPS
1. we assume that the xml input is for data record usage, that is nodes with same name should have same children. If not, we assume that it is because children with no value is hidden. so the final data structure for a node contains every children ever appeard in one of the instance.
2. there's some testing code in protofactory, importing workspace. those testing data is ignored by git. If you clone this repository from github and meet some import issue around this, simply comment out the import sentence and corresponding testing function will just work.