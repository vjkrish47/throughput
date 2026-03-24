# throughput
**Throughput Aware Replica Selection for Multi Region Distributed Systems**

* Author: Vijaya Krishna Namala
* Published In : International Journal on Science and Technology (IJSAT) 
* Publication Date: June 2024
* E-ISSN:  2229-7677
* Impact Factor : 9.88

**Abstract:**
Multi region distributed systems rely on replica selection for request routing, but static approaches ignore differences in replica processing capacity, leading to inefficient resource utilization and reduced throughput. This work introduces a throughput aware replica selection approach that aligns request distribution with runtime processing capability. By directing traffic toward higher capacity replicas and avoiding overloaded ones, the system improves overall efficiency and scalability. Experimental evaluation across increasing region counts demonstrates significantly higher throughput and better resource utilization compared to baseline methods.

**Key Contributions**

* **Throughput Aware Replica Selection Design:**\
Developed a routing mechanism that considers runtime processing capacity of replicas instead of uniform distribution to improve request handling efficiency.

* **apacity Driven Request Distribution Mechanism:**\
Introduced a selection approach that dynamically directs more requests toward higher throughput replicas while avoiding overloaded or constrained regions.

* **Baseline vs Throughput Aware Comparative Analysis:**\
Conducted detailed experiments comparing static replica selection with throughput aware routing across multiple region configurations.

* **Scalability Evaluation Across Region Sizes:**\
Analyzed system performance across 3, 5, 7, 9, and 11 regions to study throughput growth and resource utilization under increasing scale.

**Relevance & Real World Impact**

* **Improved Throughput Performance :**\
Throughput aware routing significantly increases operations per second by aligning request distribution with replica processing capability.

* **Efficient Resource Utilization :**\
The approach ensures that higher capacity replicas are effectively utilized while preventing overload on weaker replicas.

* **Enhanced Scalability in Multi Region Systems :**\
System throughput continues to grow consistently as region count increases, avoiding early saturation seen in static routing approaches.

* **Reduced Bottleneck Effects :**\
By preventing weaker replicas from limiting performance, the system achieves balanced load distribution and stable operation.

* **Academic and Practical Contribution :**\
Provides a structured foundation for capacity aware routing strategies, supporting research and real world deployment in large scale distributed systems.

**Experimental Results (Summary)**:

  | Nodes | Baseline (ops/sec) | Throughput Aware (ops/sec)| Improvement (%) |
  |-------|--------------------| --------------------------| ----------------|
  | 3     | 1800               | 2450                      | 36.11           |
  | 5     | 2300               | 3200                      | 39.13           |
  | 7     | 2700               | 3850                      | 42.59           |
  | 9     | 3050               | 4400                      | 44.26           |
  | 11    | 3350               | 4900                      | 46.27           |

**Citation** \
Throughput Aware Replica Selection for Multi Region Distributed Systems. \
Vijaya Krishna Namala \
International Journal on Science and Technology.\
E-ISSN-  2229-7677 \
**License** \
This research is shared for a academic and research purposes. For commercial use, please contact the author.\
**Resources** \
https://www.ijsat.org/ \
**Author Contact** \
**LinkedIn**: linkedin.com/in/vijaya-krishna-namala-a42b2958 | **Email**: vijaya.namala@gmail.com



