<!--
 * @Description: 
 * @Author: neozhang
 * @Date: 2022-02-10 13:04:38
 * @LastEditors: neozhang
 * @LastEditTime: 2022-02-10 13:08:00
-->
# jager原理  

微服务链路追踪作用：如果你不能度量它，你就无法改进它。  

## jager作用  

它是用来监视和诊断基于微服务的分布式系统  

- 用于服务依赖性分析，辅助性能优化  

## 主要特性  

- 高扩展性  
- 原生支持OpenTracing  
- 可观察性  

## 术语  

### Span  

- Jager中的逻辑工作单元  