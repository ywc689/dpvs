project('dpvs', 'c',
  version : '1.10.0',
#default_options : ['c_std=c99']
  )

cflags = []
sources = []
#extra_ldlibs = ['-lnuma']
extra_ldlibs = []

inc = include_directories('../include', '/usr/local/ssl/include/')

cflags += ['-D__DPVS__', '-DALLOW_EXPERIMENTAL_API', '-DDPVS_MAX_SOCKET=2', '-DDPVS_MAX_LCORE=64', '-DDPVS_VERSION="1.10.0"', '-DDPVS_BUILD_DATE="2021-01-01"', '-DCONFIG_DPVS_FDIR=n', '-Wall', '-Werror', '-Wno-pedantic', '-Wno-sign-compare', '-Wstrict-prototypes', '-Wmissing-prototypes', '-mcmodel=medium']

cflags += ['-g', '-O0', '-D', 'DEBUG']

subdir('ipvs')
subdir('ipv6')
subdir('tc')

sources += files(
		   '/usr/local/include/dpdk/rte_config.h',
           'parser.c',
           'inet.c',
           'timer.c',
           'ctrl.c',
           'route.c',
           'kni.c',
           'iftraf.c',
           'mbuf.c',
           'vlan.c',
           'ipip.c',
           'global_data.c',
           'neigh.c',
           'ip_tunnel.c',
           'pidfile.c',
           'sa_pool.c',
           'log.c',
           'ipv4.c',
           'netif_addr.c',
           'netif_flow.c',
           'sys_time.c',
           'icmp.c',
           'scheduler.c',
           'ipv4_frag.c',
           'ip_gre.c',
           'ipset.c',
           'netif.c',
           'inetaddr.c',
           'eal_mem.c',
           'pdump.c',
           'main.c',
           'common.c',
           'debug.c',
           'global_conf.c',
           'cfgfile.c',
           'mempool.c',
)

#add_project_link_arguments('-Wl,--as-needed', language: 'c')
#add_project_link_arguments('/home/huangyichen/dpdk-stable-20.11.1/patched-debug/lib64/librte_net_bond.so', language: 'c')
add_project_link_arguments('-Wl,--whole-archive', language: 'c')
#add_project_link_arguments('-Wl,-rpath=/home/huangyichen/dpdk-stable-20.11.1-latest/patched-libs/lib64/', language: 'c')
add_project_link_arguments('-L/home/huangyichen/dpdk-stable-20.11.1-latest/patched-libs/lib64/', language: 'c')
add_project_link_arguments('-l:librte_common_cpt.a', language: 'c') 
add_project_link_arguments('-l:librte_common_dpaax.a', language: 'c') 
add_project_link_arguments('-l:librte_common_iavf.a', language: 'c') 
add_project_link_arguments('-l:librte_common_octeontx.a', language: 'c') 
add_project_link_arguments('-l:librte_common_octeontx2.a', language: 'c') 
add_project_link_arguments('-l:librte_common_sfc_efx.a', language: 'c') 
add_project_link_arguments('-l:librte_bus_dpaa.a', language: 'c') 
add_project_link_arguments('-l:librte_bus_fslmc.a', language: 'c') 
add_project_link_arguments('-l:librte_bus_ifpga.a', language: 'c') 
add_project_link_arguments('-l:librte_bus_pci.a', language: 'c') 
add_project_link_arguments('-l:librte_bus_vdev.a', language: 'c') 
add_project_link_arguments('-l:librte_bus_vmbus.a', language: 'c') 
add_project_link_arguments('-l:librte_common_qat.a', language: 'c') 
add_project_link_arguments('-l:librte_mempool_bucket.a', language: 'c') 
add_project_link_arguments('-l:librte_mempool_dpaa.a', language: 'c') 
add_project_link_arguments('-l:librte_mempool_dpaa2.a', language: 'c') 
add_project_link_arguments('-l:librte_mempool_octeontx.a', language: 'c') 
add_project_link_arguments('-l:librte_mempool_octeontx2.a', language: 'c') 
add_project_link_arguments('-l:librte_mempool_ring.a', language: 'c') 
add_project_link_arguments('-l:librte_mempool_stack.a', language: 'c') 
add_project_link_arguments('-l:librte_net_af_packet.a', language: 'c') 
add_project_link_arguments('-l:librte_net_ark.a', language: 'c') 
add_project_link_arguments('-l:librte_net_atlantic.a', language: 'c') 
add_project_link_arguments('-l:librte_net_avp.a', language: 'c') 
add_project_link_arguments('-l:librte_net_axgbe.a', language: 'c') 
add_project_link_arguments('-l:librte_net_bond.a', language: 'c') 
add_project_link_arguments('-l:librte_net_bnx2x.a', language: 'c') 
add_project_link_arguments('-l:librte_net_bnxt.a', language: 'c') 
add_project_link_arguments('-l:librte_net_cxgbe.a', language: 'c') 
add_project_link_arguments('-l:librte_net_dpaa.a', language: 'c') 
add_project_link_arguments('-l:librte_net_dpaa2.a', language: 'c') 
add_project_link_arguments('-l:librte_net_e1000.a', language: 'c') 
add_project_link_arguments('-l:librte_net_ena.a', language: 'c') 
add_project_link_arguments('-l:librte_net_enetc.a', language: 'c') 
add_project_link_arguments('-l:librte_net_enic.a', language: 'c') 
add_project_link_arguments('-l:librte_net_failsafe.a', language: 'c') 
add_project_link_arguments('-l:librte_net_fm10k.a', language: 'c') 
add_project_link_arguments('-l:librte_net_i40e.a', language: 'c') 
add_project_link_arguments('-l:librte_net_hinic.a', language: 'c') 
add_project_link_arguments('-l:librte_net_hns3.a', language: 'c') 
add_project_link_arguments('-l:librte_net_iavf.a', language: 'c') 
add_project_link_arguments('-l:librte_net_ice.a', language: 'c') 
add_project_link_arguments('-l:librte_net_igc.a', language: 'c') 
add_project_link_arguments('-l:librte_net_ixgbe.a', language: 'c') 
add_project_link_arguments('-l:librte_net_kni.a', language: 'c') 
add_project_link_arguments('-l:librte_net_liquidio.a', language: 'c') 
add_project_link_arguments('-l:librte_net_memif.a', language: 'c') 
add_project_link_arguments('-l:librte_net_netvsc.a', language: 'c') 
add_project_link_arguments('-l:librte_net_nfp.a', language: 'c') 
add_project_link_arguments('-l:librte_net_null.a', language: 'c') 
add_project_link_arguments('-l:librte_net_octeontx.a', language: 'c') 
add_project_link_arguments('-l:librte_net_octeontx2.a', language: 'c') 
add_project_link_arguments('-l:librte_net_pcap.a', language: 'c') 
add_project_link_arguments('-l:librte_net_pfe.a', language: 'c') 
add_project_link_arguments('-l:librte_net_qede.a', language: 'c') 
add_project_link_arguments('-l:librte_net_ring.a', language: 'c') 
add_project_link_arguments('-l:librte_net_sfc.a', language: 'c') 
add_project_link_arguments('-l:librte_net_softnic.a', language: 'c') 
add_project_link_arguments('-l:librte_net_tap.a', language: 'c') 
add_project_link_arguments('-l:librte_net_thunderx.a', language: 'c') 
add_project_link_arguments('-l:librte_net_txgbe.a', language: 'c') 
add_project_link_arguments('-l:librte_net_vdev_netvsc.a', language: 'c') 
add_project_link_arguments('-l:librte_net_vhost.a', language: 'c') 
add_project_link_arguments('-l:librte_net_virtio.a', language: 'c') 
add_project_link_arguments('-l:librte_net_vmxnet3.a', language: 'c') 
add_project_link_arguments('-l:librte_raw_dpaa2_cmdif.a', language: 'c') 
add_project_link_arguments('-l:librte_raw_dpaa2_qdma.a', language: 'c') 
add_project_link_arguments('-l:librte_raw_ioat.a', language: 'c') 
add_project_link_arguments('-l:librte_raw_ntb.a', language: 'c') 
add_project_link_arguments('-l:librte_raw_octeontx2_dma.a', language: 'c') 
add_project_link_arguments('-l:librte_raw_octeontx2_ep.a', language: 'c') 
add_project_link_arguments('-l:librte_raw_skeleton.a', language: 'c') 
add_project_link_arguments('-l:librte_crypto_bcmfs.a', language: 'c') 
add_project_link_arguments('-l:librte_crypto_caam_jr.a', language: 'c') 
add_project_link_arguments('-l:librte_crypto_ccp.a', language: 'c') 
add_project_link_arguments('-l:librte_crypto_dpaa_sec.a', language: 'c') 
add_project_link_arguments('-l:librte_crypto_dpaa2_sec.a', language: 'c') 
add_project_link_arguments('-l:librte_crypto_nitrox.a', language: 'c') 
add_project_link_arguments('-l:librte_crypto_null.a', language: 'c') 
add_project_link_arguments('-l:librte_crypto_octeontx.a', language: 'c') 
add_project_link_arguments('-l:librte_crypto_octeontx2.a', language: 'c') 
add_project_link_arguments('-l:librte_crypto_openssl.a', language: 'c') 
add_project_link_arguments('-l:librte_crypto_scheduler.a', language: 'c') 
add_project_link_arguments('-l:librte_crypto_virtio.a', language: 'c') 
add_project_link_arguments('-l:librte_compress_octeontx.a', language: 'c') 
add_project_link_arguments('-l:librte_compress_zlib.a', language: 'c') 
add_project_link_arguments('-l:librte_regex_octeontx2.a', language: 'c') 
add_project_link_arguments('-l:librte_vdpa_ifc.a', language: 'c') 
add_project_link_arguments('-l:librte_event_dlb.a', language: 'c') 
add_project_link_arguments('-l:librte_event_dlb2.a', language: 'c') 
add_project_link_arguments('-l:librte_event_dpaa.a', language: 'c') 
add_project_link_arguments('-l:librte_event_dpaa2.a', language: 'c') 
add_project_link_arguments('-l:librte_event_octeontx2.a', language: 'c') 
add_project_link_arguments('-l:librte_event_opdl.a', language: 'c') 
add_project_link_arguments('-l:librte_event_skeleton.a', language: 'c') 
add_project_link_arguments('-l:librte_event_sw.a', language: 'c') 
add_project_link_arguments('-l:librte_event_dsw.a', language: 'c') 
add_project_link_arguments('-l:librte_event_octeontx.a', language: 'c') 
add_project_link_arguments('-l:librte_baseband_null.a', language: 'c') 
add_project_link_arguments('-l:librte_baseband_turbo_sw.a', language: 'c') 
add_project_link_arguments('-l:librte_baseband_fpga_lte_fec.a', language: 'c') 
add_project_link_arguments('-l:librte_baseband_fpga_5gnr_fec.a', language: 'c') 
add_project_link_arguments('-l:librte_baseband_acc100.a', language: 'c') 
add_project_link_arguments('-l:librte_node.a', language: 'c') 
add_project_link_arguments('-l:librte_graph.a', language: 'c') 
add_project_link_arguments('-l:librte_bpf.a', language: 'c') 
add_project_link_arguments('-l:librte_flow_classify.a', language: 'c') 
add_project_link_arguments('-l:librte_pipeline.a', language: 'c') 
add_project_link_arguments('-l:librte_table.a', language: 'c') 
add_project_link_arguments('-l:librte_port.a', language: 'c') 
add_project_link_arguments('-l:librte_fib.a', language: 'c') 
add_project_link_arguments('-l:librte_ipsec.a', language: 'c') 
add_project_link_arguments('-l:librte_vhost.a', language: 'c') 
add_project_link_arguments('-l:librte_stack.a', language: 'c') 
add_project_link_arguments('-l:librte_security.a', language: 'c') 
add_project_link_arguments('-l:librte_sched.a', language: 'c') 
add_project_link_arguments('-l:librte_reorder.a', language: 'c') 
add_project_link_arguments('-l:librte_rib.a', language: 'c') 
add_project_link_arguments('-l:librte_regexdev.a', language: 'c') 
add_project_link_arguments('-l:librte_rawdev.a', language: 'c') 
add_project_link_arguments('-l:librte_pdump.a', language: 'c') 
add_project_link_arguments('-l:librte_power.a', language: 'c') 
add_project_link_arguments('-l:librte_member.a', language: 'c') 
add_project_link_arguments('-l:librte_lpm.a', language: 'c') 
add_project_link_arguments('-l:librte_latencystats.a', language: 'c') 
add_project_link_arguments('-l:librte_kni.a', language: 'c') 
add_project_link_arguments('-l:librte_jobstats.a', language: 'c') 
add_project_link_arguments('-l:librte_ip_frag.a', language: 'c') 
add_project_link_arguments('-l:librte_gso.a', language: 'c') 
add_project_link_arguments('-l:librte_gro.a', language: 'c') 
add_project_link_arguments('-l:librte_eventdev.a', language: 'c') 
add_project_link_arguments('-l:librte_efd.a', language: 'c') 
add_project_link_arguments('-l:librte_distributor.a', language: 'c') 
add_project_link_arguments('-l:librte_cryptodev.a', language: 'c') 
add_project_link_arguments('-l:librte_compressdev.a', language: 'c') 
add_project_link_arguments('-l:librte_cfgfile.a', language: 'c') 
add_project_link_arguments('-l:librte_bitratestats.a', language: 'c') 
add_project_link_arguments('-l:librte_bbdev.a', language: 'c') 
add_project_link_arguments('-l:librte_acl.a', language: 'c') 
add_project_link_arguments('-l:librte_timer.a', language: 'c') 
add_project_link_arguments('-l:librte_hash.a', language: 'c') 
add_project_link_arguments('-l:librte_metrics.a', language: 'c') 
add_project_link_arguments('-l:librte_cmdline.a', language: 'c') 
add_project_link_arguments('-l:librte_pci.a', language: 'c') 
add_project_link_arguments('-l:librte_ethdev.a', language: 'c') 
add_project_link_arguments('-l:librte_meter.a', language: 'c') 
add_project_link_arguments('-l:librte_net.a', language: 'c') 
add_project_link_arguments('-l:librte_mbuf.a', language: 'c') 
add_project_link_arguments('-l:librte_mempool.a', language: 'c') 
add_project_link_arguments('-l:librte_rcu.a', language: 'c') 
add_project_link_arguments('-l:librte_ring.a', language: 'c') 
add_project_link_arguments('-l:librte_eal.a', language: 'c') 
add_project_link_arguments('-l:librte_telemetry.a', language: 'c') 
add_project_link_arguments('-l:librte_kvargs.a', language: 'c') 
add_project_link_arguments('-Wl,--no-whole-archive', language: 'c')
add_project_link_arguments('-Wl,--export-dynamic', language: 'c')
add_project_link_arguments('-lpcap', language: 'c')
add_project_link_arguments('-lz', language: 'c')
add_project_link_arguments('-lnuma', language: 'c')
add_project_link_arguments('-lpthread', language: 'c')
add_project_link_arguments('-lcrypto', language: 'c')
add_project_link_arguments('-lm', language: 'c')
add_project_link_arguments('-ldl', language: 'c')


executable('dpvs',
           sources,
		   dependencies: dependency('libdpdk'),
           include_directories: inc,
		   c_args: cflags,
           link_args: extra_ldlibs,
		   )
