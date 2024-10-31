/* Generated by the protocol buffer compiler.  DO NOT EDIT! */
/* Generated from: api/cluster/cluster.proto */

/* Do not generate deprecated warnings for self */
#ifndef PROTOBUF_C__NO_DEPRECATED
#define PROTOBUF_C__NO_DEPRECATED
#endif

#include "cluster/cluster.pb-c.h"
void   cluster__cluster__init
                     (Cluster__Cluster         *message)
{
  static const Cluster__Cluster init_value = CLUSTER__CLUSTER__INIT;
  *message = init_value;
}
size_t cluster__cluster__get_packed_size
                     (const Cluster__Cluster *message)
{
  assert(message->base.descriptor == &cluster__cluster__descriptor);
  return protobuf_c_message_get_packed_size ((const ProtobufCMessage*)(message));
}
size_t cluster__cluster__pack
                     (const Cluster__Cluster *message,
                      uint8_t       *out)
{
  assert(message->base.descriptor == &cluster__cluster__descriptor);
  return protobuf_c_message_pack ((const ProtobufCMessage*)message, out);
}
size_t cluster__cluster__pack_to_buffer
                     (const Cluster__Cluster *message,
                      ProtobufCBuffer *buffer)
{
  assert(message->base.descriptor == &cluster__cluster__descriptor);
  return protobuf_c_message_pack_to_buffer ((const ProtobufCMessage*)message, buffer);
}
Cluster__Cluster *
       cluster__cluster__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data)
{
  return (Cluster__Cluster *)
     protobuf_c_message_unpack (&cluster__cluster__descriptor,
                                allocator, len, data);
}
void   cluster__cluster__free_unpacked
                     (Cluster__Cluster *message,
                      ProtobufCAllocator *allocator)
{
  if(!message)
    return;
  assert(message->base.descriptor == &cluster__cluster__descriptor);
  protobuf_c_message_free_unpacked ((ProtobufCMessage*)message, allocator);
}
static const ProtobufCEnumValue cluster__cluster__lb_policy__enum_values_by_number[3] =
{
  { "ROUND_ROBIN", "CLUSTER__CLUSTER__LB_POLICY__ROUND_ROBIN", 0 },
  { "LEAST_REQUEST", "CLUSTER__CLUSTER__LB_POLICY__LEAST_REQUEST", 1 },
  { "RANDOM", "CLUSTER__CLUSTER__LB_POLICY__RANDOM", 3 },
};
static const ProtobufCIntRange cluster__cluster__lb_policy__value_ranges[] = {
{0, 0},{3, 2},{0, 3}
};
static const ProtobufCEnumValueIndex cluster__cluster__lb_policy__enum_values_by_name[3] =
{
  { "LEAST_REQUEST", 1 },
  { "RANDOM", 2 },
  { "ROUND_ROBIN", 0 },
};
const ProtobufCEnumDescriptor cluster__cluster__lb_policy__descriptor =
{
  PROTOBUF_C__ENUM_DESCRIPTOR_MAGIC,
  "cluster.Cluster.LbPolicy",
  "LbPolicy",
  "Cluster__Cluster__LbPolicy",
  "cluster",
  3,
  cluster__cluster__lb_policy__enum_values_by_number,
  3,
  cluster__cluster__lb_policy__enum_values_by_name,
  2,
  cluster__cluster__lb_policy__value_ranges,
  NULL,NULL,NULL,NULL   /* reserved[1234] */
};
static const ProtobufCFieldDescriptor cluster__cluster__field_descriptors[7] =
{
  {
    "name",
    1,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_STRING,
    0,   /* quantifier_offset */
    offsetof(Cluster__Cluster, name),
    NULL,
    &protobuf_c_empty_string,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "id",
    2,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_UINT32,
    0,   /* quantifier_offset */
    offsetof(Cluster__Cluster, id),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "connect_timeout",
    4,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_UINT32,
    0,   /* quantifier_offset */
    offsetof(Cluster__Cluster, connect_timeout),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "lb_policy",
    6,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_ENUM,
    0,   /* quantifier_offset */
    offsetof(Cluster__Cluster, lb_policy),
    &cluster__cluster__lb_policy__descriptor,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "circuit_breakers",
    10,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_MESSAGE,
    0,   /* quantifier_offset */
    offsetof(Cluster__Cluster, circuit_breakers),
    &cluster__circuit_breakers__descriptor,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "load_assignment",
    33,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_MESSAGE,
    0,   /* quantifier_offset */
    offsetof(Cluster__Cluster, load_assignment),
    &endpoint__cluster_load_assignment__descriptor,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "api_status",
    128,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_ENUM,
    0,   /* quantifier_offset */
    offsetof(Cluster__Cluster, api_status),
    &core__api_status__descriptor,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
};
static const unsigned cluster__cluster__field_indices_by_name[] = {
  6,   /* field[6] = api_status */
  4,   /* field[4] = circuit_breakers */
  2,   /* field[2] = connect_timeout */
  1,   /* field[1] = id */
  3,   /* field[3] = lb_policy */
  5,   /* field[5] = load_assignment */
  0,   /* field[0] = name */
};
static const ProtobufCIntRange cluster__cluster__number_ranges[6 + 1] =
{
  { 1, 0 },
  { 4, 2 },
  { 6, 3 },
  { 10, 4 },
  { 33, 5 },
  { 128, 6 },
  { 0, 7 }
};
const ProtobufCMessageDescriptor cluster__cluster__descriptor =
{
  PROTOBUF_C__MESSAGE_DESCRIPTOR_MAGIC,
  "cluster.Cluster",
  "Cluster",
  "Cluster__Cluster",
  "cluster",
  sizeof(Cluster__Cluster),
  7,
  cluster__cluster__field_descriptors,
  cluster__cluster__field_indices_by_name,
  6,  cluster__cluster__number_ranges,
  (ProtobufCMessageInit) cluster__cluster__init,
  NULL,NULL,NULL    /* reserved[123] */
};
