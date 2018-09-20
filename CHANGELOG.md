# Change Log

## [0.9.0-beta.0](https://github.com/kubedb/cli/tree/0.9.0-beta.0) (2018-09-20)
[Full Changelog](https://github.com/kubedb/cli/compare/0.8.0...0.9.0-beta.0)

**Fixed bugs:**

- Fix error when nothing is patched [\#294](https://github.com/kubedb/cli/pull/294) ([tamalsaha](https://github.com/tamalsaha))
- Only waity for pods in operator namespace during purging [\#293](https://github.com/kubedb/cli/pull/293) ([tamalsaha](https://github.com/tamalsaha))
- Fixed variable expression of 'KUBEDB\_CATALOG' [\#260](https://github.com/kubedb/cli/pull/260) ([the-redback](https://github.com/the-redback))
- Set namespace only for DB crds in --purge [\#258](https://github.com/kubedb/cli/pull/258) ([the-redback](https://github.com/the-redback))

**Merged pull requests:**

- Prepare docs for 0.9.0-beta.0 release [\#297](https://github.com/kubedb/cli/pull/297) ([tamalsaha](https://github.com/tamalsaha))
- Updated database versions of example CRDs to non-deprecated [\#296](https://github.com/kubedb/cli/pull/296) ([the-redback](https://github.com/the-redback))
- Release 0.9.0-beta.0 [\#295](https://github.com/kubedb/cli/pull/295) ([tamalsaha](https://github.com/tamalsaha))
- Add clusterrole for 'pod/exec' [\#292](https://github.com/kubedb/cli/pull/292) ([shudipta](https://github.com/shudipta))
- Use suffix for updated DBImage [\#291](https://github.com/kubedb/cli/pull/291) ([the-redback](https://github.com/the-redback))
- Revendor k8s.io/kubernetes [\#290](https://github.com/kubedb/cli/pull/290) ([tamalsaha](https://github.com/tamalsaha))
- Update vendor list [\#289](https://github.com/kubedb/cli/pull/289) ([tamalsaha](https://github.com/tamalsaha))
- Revendor kubernetes-1.11.3 [\#288](https://github.com/kubedb/cli/pull/288) ([tamalsaha](https://github.com/tamalsaha))
- Automate updating dependency to kubernetes-1.11.3 [\#287](https://github.com/kubedb/cli/pull/287) ([tamalsaha](https://github.com/tamalsaha))
- Support pod annotations in chart [\#286](https://github.com/kubedb/cli/pull/286) ([tamalsaha](https://github.com/tamalsaha))
- Set serviceAccount for clearner job [\#285](https://github.com/kubedb/cli/pull/285) ([tamalsaha](https://github.com/tamalsaha))
- Fixed typos [\#284](https://github.com/kubedb/cli/pull/284) ([endrec](https://github.com/endrec))
- Revendor api [\#283](https://github.com/kubedb/cli/pull/283) ([tamalsaha](https://github.com/tamalsaha))
- Cleanup webhooks when chart is deleted [\#282](https://github.com/kubedb/cli/pull/282) ([tamalsaha](https://github.com/tamalsaha))
- Use IntHash as status.observedGeneration [\#281](https://github.com/kubedb/cli/pull/281) ([tamalsaha](https://github.com/tamalsaha))
- Added etcd catalog installation line on kubedb.sh  [\#279](https://github.com/kubedb/cli/pull/279) ([the-redback](https://github.com/the-redback))
- use officially suggested exporter image [\#278](https://github.com/kubedb/cli/pull/278) ([the-redback](https://github.com/the-redback))
- Use pgadmin4:latest [\#277](https://github.com/kubedb/cli/pull/277) ([tamalsaha](https://github.com/tamalsaha))
- Update kubedb-catalog for Elasticsearch & Postgres [\#276](https://github.com/kubedb/cli/pull/276) ([hossainemruz](https://github.com/hossainemruz))
- Support separate image instead operator for mongodb exporter [\#275](https://github.com/kubedb/cli/pull/275) ([the-redback](https://github.com/the-redback))
- Add `describe etcd` command [\#274](https://github.com/kubedb/cli/pull/274) ([tamalsaha](https://github.com/tamalsaha))
- Print StorageType [\#273](https://github.com/kubedb/cli/pull/273) ([tamalsaha](https://github.com/tamalsaha))
- Print labels and annotations in describe commands [\#272](https://github.com/kubedb/cli/pull/272) ([tamalsaha](https://github.com/tamalsaha))
- Rename kubedb\_describer file [\#271](https://github.com/kubedb/cli/pull/271) ([tamalsaha](https://github.com/tamalsaha))
- Updated memcached catalog for exporter [\#270](https://github.com/kubedb/cli/pull/270) ([the-redback](https://github.com/the-redback))
- Updated mysql catalog for exporter [\#269](https://github.com/kubedb/cli/pull/269) ([the-redback](https://github.com/the-redback))
- Added Spec.version field on DBVersion CRDs [\#268](https://github.com/kubedb/cli/pull/268) ([the-redback](https://github.com/the-redback))
- Etcd catalog file added [\#267](https://github.com/kubedb/cli/pull/267) ([sanjid133](https://github.com/sanjid133))
- Improve Helm chart options [\#266](https://github.com/kubedb/cli/pull/266) ([tamalsaha](https://github.com/tamalsaha))
- Add documentation explaining how to use Kibana and x-pack monitoring with KubeDB Elasticsearch [\#265](https://github.com/kubedb/cli/pull/265) ([hossainemruz](https://github.com/hossainemruz))
- Set vendored github.com/json-iterator/go version to 1.1.5 [\#264](https://github.com/kubedb/cli/pull/264) ([tamalsaha](https://github.com/tamalsaha))
- Support adding new dependencing via revendor.py script [\#263](https://github.com/kubedb/cli/pull/263) ([tamalsaha](https://github.com/tamalsaha))
- Add revendor script [\#262](https://github.com/kubedb/cli/pull/262) ([tamalsaha](https://github.com/tamalsaha))
- Fix spelling [\#261](https://github.com/kubedb/cli/pull/261) ([tamalsaha](https://github.com/tamalsaha))
- Added DBVersions default objects && Removed '--docker-registry' and '--exporter-tag' from operator flags [\#259](https://github.com/kubedb/cli/pull/259) ([the-redback](https://github.com/the-redback))
- Backup and delete DBversions in --purge flag [\#257](https://github.com/kubedb/cli/pull/257) ([the-redback](https://github.com/the-redback))
- Add kubedb catalog [\#256](https://github.com/kubedb/cli/pull/256) ([annymsMthd](https://github.com/annymsMthd))
- Update pgadmin to official image [\#254](https://github.com/kubedb/cli/pull/254) ([tamalsaha](https://github.com/tamalsaha))
- Add Mongodb replicaset doc [\#253](https://github.com/kubedb/cli/pull/253) ([the-redback](https://github.com/the-redback))
- Removed `In RBAC \*not\* enabled cluster` section from monitoring docs [\#252](https://github.com/kubedb/cli/pull/252) ([the-redback](https://github.com/the-redback))
- Add throughput graph [\#251](https://github.com/kubedb/cli/pull/251) ([tamalsaha](https://github.com/tamalsaha))
- Don't register mutating webhooks for  DELETE op [\#250](https://github.com/kubedb/cli/pull/250) ([tamalsaha](https://github.com/tamalsaha))
- Add ConfigMap lists in Rbac [\#249](https://github.com/kubedb/cli/pull/249) ([the-redback](https://github.com/the-redback))
- Deploy in kube-system namespace using Helm [\#248](https://github.com/kubedb/cli/pull/248) ([tamalsaha](https://github.com/tamalsaha))
- Update client-go to v8.0.0 [\#247](https://github.com/kubedb/cli/pull/247) ([tamalsaha](https://github.com/tamalsaha))
- Enable status subresource for crds [\#246](https://github.com/kubedb/cli/pull/246) ([tamalsaha](https://github.com/tamalsaha))
- Format shell script [\#245](https://github.com/kubedb/cli/pull/245) ([tamalsaha](https://github.com/tamalsaha))
- Support custom configuration file [\#244](https://github.com/kubedb/cli/pull/244) ([hossainemruz](https://github.com/hossainemruz))
- Fix chart admission webhook configuration. [\#243](https://github.com/kubedb/cli/pull/243) ([tamalsaha](https://github.com/tamalsaha))
- Add etcd webhooks to installer [\#242](https://github.com/kubedb/cli/pull/242) ([sanjid133](https://github.com/sanjid133))
- Use liveness probe [\#241](https://github.com/kubedb/cli/pull/241) ([the-redback](https://github.com/the-redback))
- Support ENV variables in CRDs [\#240](https://github.com/kubedb/cli/pull/240) ([hossainemruz](https://github.com/hossainemruz))
- Upgrade pg admin [\#239](https://github.com/kubedb/cli/pull/239) ([LinAnt](https://github.com/LinAnt))

## [0.8.0](https://github.com/kubedb/cli/tree/0.8.0) (2018-06-12)
[Full Changelog](https://github.com/kubedb/cli/compare/0.8.0-rc.0...0.8.0)

**Merged pull requests:**

- Update db support info [\#238](https://github.com/kubedb/cli/pull/238) ([tamalsaha](https://github.com/tamalsaha))
- Updated doc [\#237](https://github.com/kubedb/cli/pull/237) ([the-redback](https://github.com/the-redback))
- Prepare docs for 0.8.0 release [\#236](https://github.com/kubedb/cli/pull/236) ([tamalsaha](https://github.com/tamalsaha))
- Prepare release script for 0.8.0 release [\#235](https://github.com/kubedb/cli/pull/235) ([tamalsaha](https://github.com/tamalsaha))
- Add togglable tabs for Installation: Script & Helm [\#234](https://github.com/kubedb/cli/pull/234) ([sajibcse68](https://github.com/sajibcse68))
- Removed left over documents of deprecated flag '--force' [\#233](https://github.com/kubedb/cli/pull/233) ([the-redback](https://github.com/the-redback))
- Fix chart: Use resource 100m CPU to run operator pod [\#232](https://github.com/kubedb/cli/pull/232) ([the-redback](https://github.com/the-redback))
- Use resource 100m to run operator pod [\#231](https://github.com/kubedb/cli/pull/231) ([the-redback](https://github.com/the-redback))
- Improve installer [\#230](https://github.com/kubedb/cli/pull/230) ([tamalsaha](https://github.com/tamalsaha))
- Add changelog [\#229](https://github.com/kubedb/cli/pull/229) ([tamalsaha](https://github.com/tamalsaha))

## [0.8.0-rc.0](https://github.com/kubedb/cli/tree/0.8.0-rc.0) (2018-05-28)
[Full Changelog](https://github.com/kubedb/cli/compare/0.8.0-beta.2...0.8.0-rc.0)

**Fixed bugs:**

- Update local volume info for snapshot [\#211](https://github.com/kubedb/cli/pull/211) ([tamalsaha](https://github.com/tamalsaha))

**Merged pull requests:**

- Fix page urls [\#228](https://github.com/kubedb/cli/pull/228) ([tamalsaha](https://github.com/tamalsaha))
- Fix release script [\#227](https://github.com/kubedb/cli/pull/227) ([tamalsaha](https://github.com/tamalsaha))
- Fix hugo menu [\#226](https://github.com/kubedb/cli/pull/226) ([tamalsaha](https://github.com/tamalsaha))
- Prepare 8.0.0-rc.0 [\#225](https://github.com/kubedb/cli/pull/225) ([tamalsaha](https://github.com/tamalsaha))
- Fix parent identifier of CLI docs [\#224](https://github.com/kubedb/cli/pull/224) ([the-redback](https://github.com/the-redback))
- Improve installer [\#223](https://github.com/kubedb/cli/pull/223) ([tamalsaha](https://github.com/tamalsaha))
- Clarify that user has to install bpth operator and cli [\#222](https://github.com/kubedb/cli/pull/222) ([tamalsaha](https://github.com/tamalsaha))
- Set readiness probe timeout to 5 seconds [\#221](https://github.com/kubedb/cli/pull/221) ([the-redback](https://github.com/the-redback))
-  Added plugin for Cloud Providers Auth [\#220](https://github.com/kubedb/cli/pull/220) ([the-redback](https://github.com/the-redback))
-  Support Custom Operator TAG for db-operators [\#219](https://github.com/kubedb/cli/pull/219) ([the-redback](https://github.com/the-redback))
- add documentation for separated storage [\#218](https://github.com/kubedb/cli/pull/218) ([aerokite](https://github.com/aerokite))
- Update Installer script to support dynamic operator && Skip validation while creating an CRD object [\#217](https://github.com/kubedb/cli/pull/217) ([the-redback](https://github.com/the-redback))
- Updated Docs [\#216](https://github.com/kubedb/cli/pull/216) ([the-redback](https://github.com/the-redback))
- Update client-go to 7.0.0 [\#215](https://github.com/kubedb/cli/pull/215) ([tamalsaha](https://github.com/tamalsaha))
- Delete kubedb objects after deleting operator. [\#214](https://github.com/kubedb/cli/pull/214) ([tamalsaha](https://github.com/tamalsaha))
- Fix pluralization [\#213](https://github.com/kubedb/cli/pull/213) ([tamalsaha](https://github.com/tamalsaha))
- Improve installer [\#212](https://github.com/kubedb/cli/pull/212) ([tamalsaha](https://github.com/tamalsaha))
- Add RBAC instructions for GKE cluster [\#210](https://github.com/kubedb/cli/pull/210) ([tamalsaha](https://github.com/tamalsaha))
- Update chart repository location [\#209](https://github.com/kubedb/cli/pull/209) ([tamalsaha](https://github.com/tamalsaha))
- Support installing from local installer scripts [\#208](https://github.com/kubedb/cli/pull/208) ([tamalsaha](https://github.com/tamalsaha))
- Fix install script for minikube 0.24.x \(Kube 1.8.0\) [\#207](https://github.com/kubedb/cli/pull/207) ([tamalsaha](https://github.com/tamalsaha))
- Skip downloading onessl if already exists [\#206](https://github.com/kubedb/cli/pull/206) ([tamalsaha](https://github.com/tamalsaha))
- Reorg objects deleted in uninstall command [\#205](https://github.com/kubedb/cli/pull/205) ([tamalsaha](https://github.com/tamalsaha))
- Add travis yaml [\#204](https://github.com/kubedb/cli/pull/204) ([tahsinrahman](https://github.com/tahsinrahman))
- Add --purge flag [\#200](https://github.com/kubedb/cli/pull/200) ([tamalsaha](https://github.com/tamalsaha))
- Make it clear that installer is a single command [\#199](https://github.com/kubedb/cli/pull/199) ([tamalsaha](https://github.com/tamalsaha))
- Fix installer [\#198](https://github.com/kubedb/cli/pull/198) ([tamalsaha](https://github.com/tamalsaha))
- Update chart to match RBAC best practices for charts [\#197](https://github.com/kubedb/cli/pull/197) ([tamalsaha](https://github.com/tamalsaha))
- Add checks to installer script [\#196](https://github.com/kubedb/cli/pull/196) ([tamalsaha](https://github.com/tamalsaha))
- Remove '--force' flag from CLI [\#195](https://github.com/kubedb/cli/pull/195) ([the-redback](https://github.com/the-redback))
- Added separate CLI docs for Each Individual DBs [\#194](https://github.com/kubedb/cli/pull/194) ([the-redback](https://github.com/the-redback))
- Cleanup list [\#193](https://github.com/kubedb/cli/pull/193) ([aerokite](https://github.com/aerokite))
- Fixed certificate Secret description [\#192](https://github.com/kubedb/cli/pull/192) ([aerokite](https://github.com/aerokite))
- Elasticsearch doesn't initialize from script [\#191](https://github.com/kubedb/cli/pull/191) ([aerokite](https://github.com/aerokite))

## [0.8.0-beta.2](https://github.com/kubedb/cli/tree/0.8.0-beta.2) (2018-02-28)
[Full Changelog](https://github.com/kubedb/cli/compare/0.8.0-beta.1...0.8.0-beta.2)

**Merged pull requests:**

- Fix admission webhook flag [\#190](https://github.com/kubedb/cli/pull/190) ([tamalsaha](https://github.com/tamalsaha))
- Fix chart [\#189](https://github.com/kubedb/cli/pull/189) ([tamalsaha](https://github.com/tamalsaha))
- Fix supported version [\#188](https://github.com/kubedb/cli/pull/188) ([aerokite](https://github.com/aerokite))
- Fix monitor links [\#187](https://github.com/kubedb/cli/pull/187) ([tamalsaha](https://github.com/tamalsaha))
- Fix next step links [\#186](https://github.com/kubedb/cli/pull/186) ([tamalsaha](https://github.com/tamalsaha))
- Fix doc links [\#185](https://github.com/kubedb/cli/pull/185) ([tamalsaha](https://github.com/tamalsaha))
- Fix installer scripts [\#184](https://github.com/kubedb/cli/pull/184) ([tamalsaha](https://github.com/tamalsaha))
- Prepare docs for 0.8.0-beta.2 [\#183](https://github.com/kubedb/cli/pull/183) ([tamalsaha](https://github.com/tamalsaha))
- Prepare release version 0.8.0-beta.2 [\#182](https://github.com/kubedb/cli/pull/182) ([tamalsaha](https://github.com/tamalsaha))
- Add OWNERS file to charts [\#181](https://github.com/kubedb/cli/pull/181) ([tamalsaha](https://github.com/tamalsaha))
- Various fixes [\#180](https://github.com/kubedb/cli/pull/180) ([tamalsaha](https://github.com/tamalsaha))
- List of external tools dependency [\#179](https://github.com/kubedb/cli/pull/179) ([aerokite](https://github.com/aerokite))
- Remove InitFlags [\#178](https://github.com/kubedb/cli/pull/178) ([tamalsaha](https://github.com/tamalsaha))
- update validation call [\#177](https://github.com/kubedb/cli/pull/177) ([aerokite](https://github.com/aerokite))
- Change service documentation [\#176](https://github.com/kubedb/cli/pull/176) ([aerokite](https://github.com/aerokite))
- Use updated apiserver info [\#175](https://github.com/kubedb/cli/pull/175) ([tamalsaha](https://github.com/tamalsaha))
- Fix pointer and validation [\#174](https://github.com/kubedb/cli/pull/174) ([aerokite](https://github.com/aerokite))
- Added rbac configuration in installer doc [\#173](https://github.com/kubedb/cli/pull/173) ([the-redback](https://github.com/the-redback))
-  Update Elasticsearch documentation [\#172](https://github.com/kubedb/cli/pull/172) ([aerokite](https://github.com/aerokite))
- Simplify user role definitions [\#171](https://github.com/kubedb/cli/pull/171) ([tamalsaha](https://github.com/tamalsaha))
- Create user facing aggregate role [\#170](https://github.com/kubedb/cli/pull/170) ([tamalsaha](https://github.com/tamalsaha))
- Use official code generator scripts [\#169](https://github.com/kubedb/cli/pull/169) ([tamalsaha](https://github.com/tamalsaha))
- Fix pluralization of Elasticsearch [\#168](https://github.com/kubedb/cli/pull/168) ([tamalsaha](https://github.com/tamalsaha))
- Fix Monitoring Port in kubedb describer [\#167](https://github.com/kubedb/cli/pull/167) ([the-redback](https://github.com/the-redback))
- Use separate urls for each type of KubeDB resource [\#166](https://github.com/kubedb/cli/pull/166) ([tamalsaha](https://github.com/tamalsaha))
- Add installer script [\#165](https://github.com/kubedb/cli/pull/165) ([tamalsaha](https://github.com/tamalsaha))
-  Improve KubeDB Memcached Documentation [\#164](https://github.com/kubedb/cli/pull/164) ([the-redback](https://github.com/the-redback))
-  Improve redis docs [\#162](https://github.com/kubedb/cli/pull/162) ([the-redback](https://github.com/the-redback))
- Improve Mysql docs [\#161](https://github.com/kubedb/cli/pull/161) ([the-redback](https://github.com/the-redback))
-  Update Postgres documentation [\#160](https://github.com/kubedb/cli/pull/160) ([aerokite](https://github.com/aerokite))
- Improve structure of  MongoDB User Guide [\#159](https://github.com/kubedb/cli/pull/159) ([the-redback](https://github.com/the-redback))
- Added show-secret in Describer for mysql and mongodb [\#158](https://github.com/kubedb/cli/pull/158) ([the-redback](https://github.com/the-redback))
- Update RBAC for Job watcher [\#157](https://github.com/kubedb/cli/pull/157) ([aerokite](https://github.com/aerokite))

## [0.8.0-beta.1](https://github.com/kubedb/cli/tree/0.8.0-beta.1) (2018-01-29)
[Full Changelog](https://github.com/kubedb/cli/compare/0.8.0-beta.0...0.8.0-beta.1)

**Merged pull requests:**

- Compress binaries [\#156](https://github.com/kubedb/cli/pull/156) ([tamalsaha](https://github.com/tamalsaha))
- Update release script for 0.8.0-beta.1 [\#155](https://github.com/kubedb/cli/pull/155) ([tamalsaha](https://github.com/tamalsaha))
- Update dependencies to client-go 6.0.0 [\#154](https://github.com/kubedb/cli/pull/154) ([aerokite](https://github.com/aerokite))
- Modify RBAC Role for kubedb-operator [\#152](https://github.com/kubedb/cli/pull/152) ([aerokite](https://github.com/aerokite))

## [0.8.0-beta.0](https://github.com/kubedb/cli/tree/0.8.0-beta.0) (2018-01-07)
[Full Changelog](https://github.com/kubedb/cli/compare/0.7.1...0.8.0-beta.0)

**Merged pull requests:**

- Fix docs [\#150](https://github.com/kubedb/cli/pull/150) ([tamalsaha](https://github.com/tamalsaha))
- Use release script in cli [\#149](https://github.com/kubedb/cli/pull/149) ([tamalsaha](https://github.com/tamalsaha))
- Fix release script [\#148](https://github.com/kubedb/cli/pull/148) ([tamalsaha](https://github.com/tamalsaha))
- Reorganize files & write front matter [\#147](https://github.com/kubedb/cli/pull/147) ([sajibcse68](https://github.com/sajibcse68))
- Mount imagePullSecret in operator [\#146](https://github.com/kubedb/cli/pull/146) ([aerokite](https://github.com/aerokite))
- Add --all flag for deletes in cli [\#145](https://github.com/kubedb/cli/pull/145) ([the-redback](https://github.com/the-redback))
- remove operator image registry from name [\#144](https://github.com/kubedb/cli/pull/144) ([aerokite](https://github.com/aerokite))
- Add --force delete flag to delete kubedb CRD objects [\#143](https://github.com/kubedb/cli/pull/143) ([the-redback](https://github.com/the-redback))
- Revendor individual operators [\#142](https://github.com/kubedb/cli/pull/142) ([tamalsaha](https://github.com/tamalsaha))
- Revert back 0.7.1 in readme file [\#141](https://github.com/kubedb/cli/pull/141) ([tamalsaha](https://github.com/tamalsaha))
- Update  docs for 0.8.0 version [\#140](https://github.com/kubedb/cli/pull/140) ([aerokite](https://github.com/aerokite))
- Revendor [\#139](https://github.com/kubedb/cli/pull/139) ([tamalsaha](https://github.com/tamalsaha))
- Fix RBAC for CRD and analytics [\#138](https://github.com/kubedb/cli/pull/138) ([tamalsaha](https://github.com/tamalsaha))
- Set ClientID for analytics [\#137](https://github.com/kubedb/cli/pull/137) ([tamalsaha](https://github.com/tamalsaha))
- Added more details  [\#136](https://github.com/kubedb/cli/pull/136) ([aerokite](https://github.com/aerokite))
- Add New DB documents [\#135](https://github.com/kubedb/cli/pull/135) ([the-redback](https://github.com/the-redback))
- Update pkg paths to kubedb org [\#133](https://github.com/kubedb/cli/pull/133) ([tamalsaha](https://github.com/tamalsaha))
- Add kubedb/memcached to CLI [\#132](https://github.com/kubedb/cli/pull/132) ([the-redback](https://github.com/the-redback))
- Fix aliases, modify menu-name, change left\_name -\> menu\_name [\#130](https://github.com/kubedb/cli/pull/130) ([sajibcse68](https://github.com/sajibcse68))
- Add kubedb/mongodb to cli [\#129](https://github.com/kubedb/cli/pull/129) ([the-redback](https://github.com/the-redback))
- Add kubedb/mysql to cli [\#128](https://github.com/kubedb/cli/pull/128) ([the-redback](https://github.com/the-redback))
- Add front matter for docs v0.7.1 [\#127](https://github.com/kubedb/cli/pull/127) ([sajibcse68](https://github.com/sajibcse68))
- Add front matter for cli [\#125](https://github.com/kubedb/cli/pull/125) ([tamalsaha](https://github.com/tamalsaha))
- Use prometheus tools from appscode/kutil [\#124](https://github.com/kubedb/cli/pull/124) ([tamalsaha](https://github.com/tamalsaha))
- Add quotes to jsonpath in elasticsearch tutorial [\#123](https://github.com/kubedb/cli/pull/123) ([mynameiswhm](https://github.com/mynameiswhm))
- Add missing `apiVersion: v1` for service definition [\#122](https://github.com/kubedb/cli/pull/122) ([tamalsaha](https://github.com/tamalsaha))
- Use client-go 5.x [\#121](https://github.com/kubedb/cli/pull/121) ([tamalsaha](https://github.com/tamalsaha))
- Modify switch case to sync with apiVersion [\#115](https://github.com/kubedb/cli/pull/115) ([aerokite](https://github.com/aerokite))

## [0.7.1](https://github.com/kubedb/cli/tree/0.7.1) (2017-10-04)
[Full Changelog](https://github.com/kubedb/cli/compare/0.7.0...0.7.1)

## [0.7.0](https://github.com/kubedb/cli/tree/0.7.0) (2017-09-26)
[Full Changelog](https://github.com/kubedb/cli/compare/0.6.0...0.7.0)

**Merged pull requests:**

- Update TPR tp CRD in docs [\#116](https://github.com/kubedb/cli/pull/116) ([tamalsaha](https://github.com/tamalsaha))
- Support CRD [\#114](https://github.com/kubedb/cli/pull/114) ([aerokite](https://github.com/aerokite))
- Remove \<invalid\> option [\#112](https://github.com/kubedb/cli/pull/112) ([aerokite](https://github.com/aerokite))
- Add DCO [\#111](https://github.com/kubedb/cli/pull/111) ([tamalsaha](https://github.com/tamalsaha))
- Highlight pgadmin username & password [\#110](https://github.com/kubedb/cli/pull/110) ([tamalsaha](https://github.com/tamalsaha))

## [0.6.0](https://github.com/kubedb/cli/tree/0.6.0) (2017-07-24)
[Full Changelog](https://github.com/kubedb/cli/compare/0.5.0...0.6.0)

**Merged pull requests:**

- Prepare docs for 0.6.0 release. [\#107](https://github.com/kubedb/cli/pull/107) ([tamalsaha](https://github.com/tamalsaha))

## [0.5.0](https://github.com/kubedb/cli/tree/0.5.0) (2017-07-19)
[Full Changelog](https://github.com/kubedb/cli/compare/0.4.0...0.5.0)

**Merged pull requests:**

- Fix concept docs. [\#105](https://github.com/kubedb/cli/pull/105) ([tamalsaha](https://github.com/tamalsaha))
- fix jsonpath typo [\#104](https://github.com/kubedb/cli/pull/104) ([diptadas](https://github.com/diptadas))
- Cross link docs [\#102](https://github.com/kubedb/cli/pull/102) ([tamalsaha](https://github.com/tamalsaha))
- Reorganize user guide [\#98](https://github.com/kubedb/cli/pull/98) ([tamalsaha](https://github.com/tamalsaha))
- Add tutorials [\#88](https://github.com/kubedb/cli/pull/88) ([tamalsaha](https://github.com/tamalsaha))

## [0.4.0](https://github.com/kubedb/cli/tree/0.4.0) (2017-07-18)
[Full Changelog](https://github.com/kubedb/cli/compare/0.3.1...0.4.0)

**Merged pull requests:**

- Print messages for RBAC objects created / updated [\#92](https://github.com/kubedb/cli/pull/92) ([tamalsaha](https://github.com/tamalsaha))
- Apply app=kubedb labels to RBAC objects [\#97](https://github.com/kubedb/cli/pull/97) ([tamalsaha](https://github.com/tamalsaha))
- Add storageclasses GET permission for kubedb-operator [\#96](https://github.com/kubedb/cli/pull/96) ([tamalsaha](https://github.com/tamalsaha))
- Fix RBAC roles for kubedb-operator [\#95](https://github.com/kubedb/cli/pull/95) ([tamalsaha](https://github.com/tamalsaha))
- Add validation for delete command [\#94](https://github.com/kubedb/cli/pull/94) ([aerokite](https://github.com/aerokite))
- Prepare docs for 0.4.0 release. [\#93](https://github.com/kubedb/cli/pull/93) ([tamalsaha](https://github.com/tamalsaha))

## [0.3.1](https://github.com/kubedb/cli/tree/0.3.1) (2017-07-17)
[Full Changelog](https://github.com/kubedb/cli/compare/0.3.0...0.3.1)

## [0.3.0](https://github.com/kubedb/cli/tree/0.3.0) (2017-07-08)
[Full Changelog](https://github.com/kubedb/cli/compare/0.2.0...0.3.0)

**Merged pull requests:**

- Add ./hack/release.py script [\#86](https://github.com/kubedb/cli/pull/86) ([tamalsaha](https://github.com/tamalsaha))
- Use VerbAll for KubeDB tprs [\#85](https://github.com/kubedb/cli/pull/85) ([tamalsaha](https://github.com/tamalsaha))
- Revendor apimachinery [\#84](https://github.com/kubedb/cli/pull/84) ([tamalsaha](https://github.com/tamalsaha))
- Document how to use swift [\#82](https://github.com/kubedb/cli/pull/82) ([tamalsaha](https://github.com/tamalsaha))
- Fix build [\#81](https://github.com/kubedb/cli/pull/81) ([tamalsaha](https://github.com/tamalsaha))
- Part 3 - User Guide [\#80](https://github.com/kubedb/cli/pull/80) ([tamalsaha](https://github.com/tamalsaha))
- Part 2 - User Guide [\#79](https://github.com/kubedb/cli/pull/79) ([tamalsaha](https://github.com/tamalsaha))
- Change KubeDB labels [\#78](https://github.com/kubedb/cli/pull/78) ([tamalsaha](https://github.com/tamalsaha))
- Fixed a typo in monitor docs [\#73](https://github.com/kubedb/cli/pull/73) ([farcaller](https://github.com/farcaller))
- Part 1 - User Guide [\#77](https://github.com/kubedb/cli/pull/77) ([tamalsaha](https://github.com/tamalsaha))
- Create RBAC roles during installation [\#76](https://github.com/kubedb/cli/pull/76) ([tamalsaha](https://github.com/tamalsaha))
- Update tag line [\#75](https://github.com/kubedb/cli/pull/75) ([tamalsaha](https://github.com/tamalsaha))
- Support non-default service account to install operator [\#72](https://github.com/kubedb/cli/pull/72) ([tamalsaha](https://github.com/tamalsaha))

## [0.2.0](https://github.com/kubedb/cli/tree/0.2.0) (2017-06-23)
[Full Changelog](https://github.com/kubedb/cli/compare/0.1.0...0.2.0)

**Merged pull requests:**

- Add validation [\#69](https://github.com/kubedb/cli/pull/69) ([aerokite](https://github.com/aerokite))
- Regenrate cli docs [\#68](https://github.com/kubedb/cli/pull/68) ([tamalsaha](https://github.com/tamalsaha))
- Various changes to summary report commands. [\#66](https://github.com/kubedb/cli/pull/66) ([tamalsaha](https://github.com/tamalsaha))
- Generate summary report & compare [\#65](https://github.com/kubedb/cli/pull/65) ([aerokite](https://github.com/aerokite))
- use correct ObjectReference [\#62](https://github.com/kubedb/cli/pull/62) ([aerokite](https://github.com/aerokite))
- Change operator namespace to kube-system [\#63](https://github.com/kubedb/cli/pull/63) ([tamalsaha](https://github.com/tamalsaha))
- Use client-go [\#61](https://github.com/kubedb/cli/pull/61) ([tamalsaha](https://github.com/tamalsaha))

## [0.1.0](https://github.com/kubedb/cli/tree/0.1.0) (2017-06-14)
**Fixed bugs:**

- Fix switch statements bug [\#44](https://github.com/kubedb/cli/pull/44) ([aerokite](https://github.com/aerokite))

**Merged pull requests:**

- Change api versionto v1alpha1 [\#60](https://github.com/kubedb/cli/pull/60) ([tamalsaha](https://github.com/tamalsaha))
- Show Kind & DB name in snapshot list [\#58](https://github.com/kubedb/cli/pull/58) ([aerokite](https://github.com/aerokite))
- Add User Guide [\#53](https://github.com/kubedb/cli/pull/53) ([aerokite](https://github.com/aerokite))
- Use built-in exporter [\#52](https://github.com/kubedb/cli/pull/52) ([tamalsaha](https://github.com/tamalsaha))
- Check conditional Precondition [\#51](https://github.com/kubedb/cli/pull/51) ([aerokite](https://github.com/aerokite))
- Add analytics for cli [\#49](https://github.com/kubedb/cli/pull/49) ([tamalsaha](https://github.com/tamalsaha))
- init command creates kubedb-operator service [\#48](https://github.com/kubedb/cli/pull/48) ([aerokite](https://github.com/aerokite))
- Integrate prometheus monitoring [\#46](https://github.com/kubedb/cli/pull/46) ([aerokite](https://github.com/aerokite))
- Use group name with thirdpartyresource [\#42](https://github.com/kubedb/cli/pull/42) ([aerokite](https://github.com/aerokite))
- Move main file to root folder [\#40](https://github.com/kubedb/cli/pull/40) ([tamalsaha](https://github.com/tamalsaha))
- Show snapshot list in dormantdatabase description [\#33](https://github.com/kubedb/cli/pull/33) ([aerokite](https://github.com/aerokite))
- Fix go report card issues [\#32](https://github.com/kubedb/cli/pull/32) ([tamalsaha](https://github.com/tamalsaha))
- Rename DeletedDatabase to DormantDatabase [\#31](https://github.com/kubedb/cli/pull/31) ([tamalsaha](https://github.com/tamalsaha))
- Use ResourceCode [\#23](https://github.com/kubedb/cli/pull/23) ([aerokite](https://github.com/aerokite))
- Remove describing DatabaseSecret [\#20](https://github.com/kubedb/cli/pull/20) ([aerokite](https://github.com/aerokite))
- Delete Docker files [\#19](https://github.com/kubedb/cli/pull/19) ([tamalsaha](https://github.com/tamalsaha))
- Delete elastic discovery [\#18](https://github.com/kubedb/cli/pull/18) ([tamalsaha](https://github.com/tamalsaha))
- Implement "kubedb describe" [\#17](https://github.com/kubedb/cli/pull/17) ([aerokite](https://github.com/aerokite))
- Implement "kubedb delete" [\#15](https://github.com/kubedb/cli/pull/15) ([aerokite](https://github.com/aerokite))
- Implement "kubedb create" [\#13](https://github.com/kubedb/cli/pull/13) ([aerokite](https://github.com/aerokite))
- Implement "kubedb get" [\#12](https://github.com/kubedb/cli/pull/12) ([aerokite](https://github.com/aerokite))
- Add elastic\_discovery [\#6](https://github.com/kubedb/cli/pull/6) ([tamalsaha](https://github.com/tamalsaha))
- Release database docker files. [\#5](https://github.com/kubedb/cli/pull/5) ([tamalsaha](https://github.com/tamalsaha))
- Release database docker files. [\#4](https://github.com/kubedb/cli/pull/4) ([tamalsaha](https://github.com/tamalsaha))
- Add ./hack/gendocs/make.sh command generate docs. [\#57](https://github.com/kubedb/cli/pull/57) ([tamalsaha](https://github.com/tamalsaha))
- Support passing namespace with kubedb commands [\#45](https://github.com/kubedb/cli/pull/45) ([aerokite](https://github.com/aerokite))
- Implement "kubedb init" [\#35](https://github.com/kubedb/cli/pull/35) ([aerokite](https://github.com/aerokite))
- Implement "kubedb edit" [\#34](https://github.com/kubedb/cli/pull/34) ([aerokite](https://github.com/aerokite))
- Fix nil pointer panic [\#26](https://github.com/kubedb/cli/pull/26) ([aerokite](https://github.com/aerokite))
- Remove delete --all option [\#25](https://github.com/kubedb/cli/pull/25) ([aerokite](https://github.com/aerokite))
- Rename DatabaseSnapshot to Snapshot [\#24](https://github.com/kubedb/cli/pull/24) ([tamalsaha](https://github.com/tamalsaha))



\* *This Change Log was automatically generated by [github_changelog_generator](https://github.com/skywinder/Github-Changelog-Generator)*