# storage_optimization
STORAGE OPTIMIZATION IN DISTRIBUTED ENVIRONMENTS USING OPTIMISTIC CONCURRENCY CONTROL

# thomas_write_rule
**STREAMLINING TRANSACTION COMMIT FOR DISTRIBUTED DATABASES WITH THOMAS'S WRITE RULE**
* Author: Vipul Reddy
* Published In : International Journal For Multidisciplinary Research (IJFMR)
* Publication Date: 08-2023
* E-ISSN: 2852-2160
* Impact Factor: 9.24
* Link:

**Abstract:**\
This paper addresses performance degradation in database transaction management caused by high abort rates under Basic Timestamp Ordering (BTO). It examines how strict timestamp-based conflict resolution leads to frequent transaction rollbacks, increased retries, and wasted computation in high-contention and long-running transaction environments. The study emphasizes the limitations of BTO, including rigid conflict handling, timestamp management overhead, and the lack of effective deadlock resolution in distributed systems. By incorporating Thomas’s Write Rule, the proposed approach allows obsolete write operations to be safely ignored, thereby reducing unnecessary aborts while preserving serializability. The paper highlights the need for enhanced timestamp-based concurrency control mechanisms to improve scalability, throughput, and efficiency in high-volume database systems.

**Key Contributions:**
* **Abort Rate Mitigation:**\
Investigated the limitations of Basic Timestamp Ordering that result in frequent transaction aborts under high contention and long-running workloads.

* **Timestamp-Based Optimization:**\
Applied Thomas’s Write Rule to relax strict write conflict handling, allowing obsolete writes to be ignored without violating serializability.
  
* **Experimental Assessment:** \
  Evaluated the impact of the proposed approach across varying workload intensities, demonstrating reduced aborts, fewer retries, and improved throughput.
  
* **Scholarly Leadership:**\
  Directed the analysis and implementation focused on enhancing scalability and efficiency of timestamp-based concurrency control mechanisms.

**Relevance & Real-World Impact**
* **Enhanced System Efficiency:**\
Improved transaction completion rates by minimizing unnecessary rollbacks in timestamp-ordered database systems.

* **Better Handling of High Contention:**\
Enabled more resilient concurrency management in environments with heavy write conflicts and long-running transactions.

* **Overall Performance Gains :** \
    Reduced wasted computation and execution delays, leading to lower latency and improved system throughput.
  
* **Academic and Educational Value:** \
    Provides a strong foundation for further research and teaching in concurrency control, transaction scheduling, and distributed database systems.

**Experimental Results (Summary)**:

  | Nodes | Basic Timestamp Ordering BTO | Thomas Write Rule TWR   | Reduction (%)   |
  |-------|------------------------------| ------------------------| ----------------|
  | 3     |  4                           | 2                       | 50.00           |
  | 5     |  10                          | 5                       | 50.00           |
  | 7     |  18                          | 8                       | 55.56           |
  | 9     |  27                          | 11                      | 59.26           |
  | 11    |  39                          | 14                      | 64.10           |

**Citation** \
STREAMLINING TRANSACTION COMMIT FOR DISTRIBUTED DATABASES WITH THOMAS'S WRITE RULE
* Vipul R 
* International Journal For Multidisciplinary Research 
* E-ISSN 2852-2160
* License \
This research is shared for a academic and research purposes. For commercial use, please contact the author.\
**Resources** \
https://www.ijfmr.com/ \
**Author Contact** \
**LinkedIn**: http://linkedin.com/in/Please add here | **Email**: please keep email id @gmail.com






