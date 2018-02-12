/**
 * Autogenerated by Thrift Compiler (0.9.3)
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */
import org.apache.thrift.scheme.IScheme;
import org.apache.thrift.scheme.SchemeFactory;
import org.apache.thrift.scheme.StandardScheme;

import org.apache.thrift.scheme.TupleScheme;
import org.apache.thrift.protocol.TTupleProtocol;
import org.apache.thrift.protocol.TProtocolException;
import org.apache.thrift.EncodingUtils;
import org.apache.thrift.TException;
import org.apache.thrift.async.AsyncMethodCallback;
import org.apache.thrift.server.AbstractNonblockingServer.*;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.util.HashMap;
import java.util.EnumMap;
import java.util.Set;
import java.util.HashSet;
import java.util.EnumSet;
import java.util.Collections;
import java.util.BitSet;
import java.nio.ByteBuffer;
import java.util.Arrays;
import javax.annotation.Generated;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

@SuppressWarnings({"cast", "rawtypes", "serial", "unchecked"})
@Generated(value = "Autogenerated by Thrift Compiler (0.9.3)", date = "2018-01-10")
public class TaskDefinition implements org.apache.thrift.TBase<TaskDefinition, TaskDefinition._Fields>, java.io.Serializable, Cloneable, Comparable<TaskDefinition> {
  private static final org.apache.thrift.protocol.TStruct STRUCT_DESC = new org.apache.thrift.protocol.TStruct("TaskDefinition");

  private static final org.apache.thrift.protocol.TField COMMAND_FIELD_DESC = new org.apache.thrift.protocol.TField("command", org.apache.thrift.protocol.TType.STRUCT, (short)1);
  private static final org.apache.thrift.protocol.TField SNAPSHOT_ID_FIELD_DESC = new org.apache.thrift.protocol.TField("snapshotId", org.apache.thrift.protocol.TType.STRING, (short)2);
  private static final org.apache.thrift.protocol.TField TASK_ID_FIELD_DESC = new org.apache.thrift.protocol.TField("taskId", org.apache.thrift.protocol.TType.STRING, (short)3);
  private static final org.apache.thrift.protocol.TField TIMEOUT_MS_FIELD_DESC = new org.apache.thrift.protocol.TField("timeoutMs", org.apache.thrift.protocol.TType.I32, (short)4);

  private static final Map<Class<? extends IScheme>, SchemeFactory> schemes = new HashMap<Class<? extends IScheme>, SchemeFactory>();
  static {
    schemes.put(StandardScheme.class, new TaskDefinitionStandardSchemeFactory());
    schemes.put(TupleScheme.class, new TaskDefinitionTupleSchemeFactory());
  }

  public Command command; // required
  public String snapshotId; // optional
  public String taskId; // optional
  public int timeoutMs; // optional

  /** The set of fields this struct contains, along with convenience methods for finding and manipulating them. */
  public enum _Fields implements org.apache.thrift.TFieldIdEnum {
    COMMAND((short)1, "command"),
    SNAPSHOT_ID((short)2, "snapshotId"),
    TASK_ID((short)3, "taskId"),
    TIMEOUT_MS((short)4, "timeoutMs");

    private static final Map<String, _Fields> byName = new HashMap<String, _Fields>();

    static {
      for (_Fields field : EnumSet.allOf(_Fields.class)) {
        byName.put(field.getFieldName(), field);
      }
    }

    /**
     * Find the _Fields constant that matches fieldId, or null if its not found.
     */
    public static _Fields findByThriftId(int fieldId) {
      switch(fieldId) {
        case 1: // COMMAND
          return COMMAND;
        case 2: // SNAPSHOT_ID
          return SNAPSHOT_ID;
        case 3: // TASK_ID
          return TASK_ID;
        case 4: // TIMEOUT_MS
          return TIMEOUT_MS;
        default:
          return null;
      }
    }

    /**
     * Find the _Fields constant that matches fieldId, throwing an exception
     * if it is not found.
     */
    public static _Fields findByThriftIdOrThrow(int fieldId) {
      _Fields fields = findByThriftId(fieldId);
      if (fields == null) throw new IllegalArgumentException("Field " + fieldId + " doesn't exist!");
      return fields;
    }

    /**
     * Find the _Fields constant that matches name, or null if its not found.
     */
    public static _Fields findByName(String name) {
      return byName.get(name);
    }

    private final short _thriftId;
    private final String _fieldName;

    _Fields(short thriftId, String fieldName) {
      _thriftId = thriftId;
      _fieldName = fieldName;
    }

    public short getThriftFieldId() {
      return _thriftId;
    }

    public String getFieldName() {
      return _fieldName;
    }
  }

  // isset id assignments
  private static final int __TIMEOUTMS_ISSET_ID = 0;
  private byte __isset_bitfield = 0;
  private static final _Fields optionals[] = {_Fields.SNAPSHOT_ID,_Fields.TASK_ID,_Fields.TIMEOUT_MS};
  public static final Map<_Fields, org.apache.thrift.meta_data.FieldMetaData> metaDataMap;
  static {
    Map<_Fields, org.apache.thrift.meta_data.FieldMetaData> tmpMap = new EnumMap<_Fields, org.apache.thrift.meta_data.FieldMetaData>(_Fields.class);
    tmpMap.put(_Fields.COMMAND, new org.apache.thrift.meta_data.FieldMetaData("command", org.apache.thrift.TFieldRequirementType.REQUIRED, 
        new org.apache.thrift.meta_data.StructMetaData(org.apache.thrift.protocol.TType.STRUCT, Command.class)));
    tmpMap.put(_Fields.SNAPSHOT_ID, new org.apache.thrift.meta_data.FieldMetaData("snapshotId", org.apache.thrift.TFieldRequirementType.OPTIONAL, 
        new org.apache.thrift.meta_data.FieldValueMetaData(org.apache.thrift.protocol.TType.STRING)));
    tmpMap.put(_Fields.TASK_ID, new org.apache.thrift.meta_data.FieldMetaData("taskId", org.apache.thrift.TFieldRequirementType.OPTIONAL, 
        new org.apache.thrift.meta_data.FieldValueMetaData(org.apache.thrift.protocol.TType.STRING)));
    tmpMap.put(_Fields.TIMEOUT_MS, new org.apache.thrift.meta_data.FieldMetaData("timeoutMs", org.apache.thrift.TFieldRequirementType.OPTIONAL, 
        new org.apache.thrift.meta_data.FieldValueMetaData(org.apache.thrift.protocol.TType.I32)));
    metaDataMap = Collections.unmodifiableMap(tmpMap);
    org.apache.thrift.meta_data.FieldMetaData.addStructMetaDataMap(TaskDefinition.class, metaDataMap);
  }

  public TaskDefinition() {
  }

  public TaskDefinition(
    Command command)
  {
    this();
    this.command = command;
  }

  /**
   * Performs a deep copy on <i>other</i>.
   */
  public TaskDefinition(TaskDefinition other) {
    __isset_bitfield = other.__isset_bitfield;
    if (other.isSetCommand()) {
      this.command = new Command(other.command);
    }
    if (other.isSetSnapshotId()) {
      this.snapshotId = other.snapshotId;
    }
    if (other.isSetTaskId()) {
      this.taskId = other.taskId;
    }
    this.timeoutMs = other.timeoutMs;
  }

  public TaskDefinition deepCopy() {
    return new TaskDefinition(this);
  }

  @Override
  public void clear() {
    this.command = null;
    this.snapshotId = null;
    this.taskId = null;
    setTimeoutMsIsSet(false);
    this.timeoutMs = 0;
  }

  public Command getCommand() {
    return this.command;
  }

  public TaskDefinition setCommand(Command command) {
    this.command = command;
    return this;
  }

  public void unsetCommand() {
    this.command = null;
  }

  /** Returns true if field command is set (has been assigned a value) and false otherwise */
  public boolean isSetCommand() {
    return this.command != null;
  }

  public void setCommandIsSet(boolean value) {
    if (!value) {
      this.command = null;
    }
  }

  public String getSnapshotId() {
    return this.snapshotId;
  }

  public TaskDefinition setSnapshotId(String snapshotId) {
    this.snapshotId = snapshotId;
    return this;
  }

  public void unsetSnapshotId() {
    this.snapshotId = null;
  }

  /** Returns true if field snapshotId is set (has been assigned a value) and false otherwise */
  public boolean isSetSnapshotId() {
    return this.snapshotId != null;
  }

  public void setSnapshotIdIsSet(boolean value) {
    if (!value) {
      this.snapshotId = null;
    }
  }

  public String getTaskId() {
    return this.taskId;
  }

  public TaskDefinition setTaskId(String taskId) {
    this.taskId = taskId;
    return this;
  }

  public void unsetTaskId() {
    this.taskId = null;
  }

  /** Returns true if field taskId is set (has been assigned a value) and false otherwise */
  public boolean isSetTaskId() {
    return this.taskId != null;
  }

  public void setTaskIdIsSet(boolean value) {
    if (!value) {
      this.taskId = null;
    }
  }

  public int getTimeoutMs() {
    return this.timeoutMs;
  }

  public TaskDefinition setTimeoutMs(int timeoutMs) {
    this.timeoutMs = timeoutMs;
    setTimeoutMsIsSet(true);
    return this;
  }

  public void unsetTimeoutMs() {
    __isset_bitfield = EncodingUtils.clearBit(__isset_bitfield, __TIMEOUTMS_ISSET_ID);
  }

  /** Returns true if field timeoutMs is set (has been assigned a value) and false otherwise */
  public boolean isSetTimeoutMs() {
    return EncodingUtils.testBit(__isset_bitfield, __TIMEOUTMS_ISSET_ID);
  }

  public void setTimeoutMsIsSet(boolean value) {
    __isset_bitfield = EncodingUtils.setBit(__isset_bitfield, __TIMEOUTMS_ISSET_ID, value);
  }

  public void setFieldValue(_Fields field, Object value) {
    switch (field) {
    case COMMAND:
      if (value == null) {
        unsetCommand();
      } else {
        setCommand((Command)value);
      }
      break;

    case SNAPSHOT_ID:
      if (value == null) {
        unsetSnapshotId();
      } else {
        setSnapshotId((String)value);
      }
      break;

    case TASK_ID:
      if (value == null) {
        unsetTaskId();
      } else {
        setTaskId((String)value);
      }
      break;

    case TIMEOUT_MS:
      if (value == null) {
        unsetTimeoutMs();
      } else {
        setTimeoutMs((Integer)value);
      }
      break;

    }
  }

  public Object getFieldValue(_Fields field) {
    switch (field) {
    case COMMAND:
      return getCommand();

    case SNAPSHOT_ID:
      return getSnapshotId();

    case TASK_ID:
      return getTaskId();

    case TIMEOUT_MS:
      return getTimeoutMs();

    }
    throw new IllegalStateException();
  }

  /** Returns true if field corresponding to fieldID is set (has been assigned a value) and false otherwise */
  public boolean isSet(_Fields field) {
    if (field == null) {
      throw new IllegalArgumentException();
    }

    switch (field) {
    case COMMAND:
      return isSetCommand();
    case SNAPSHOT_ID:
      return isSetSnapshotId();
    case TASK_ID:
      return isSetTaskId();
    case TIMEOUT_MS:
      return isSetTimeoutMs();
    }
    throw new IllegalStateException();
  }

  @Override
  public boolean equals(Object that) {
    if (that == null)
      return false;
    if (that instanceof TaskDefinition)
      return this.equals((TaskDefinition)that);
    return false;
  }

  public boolean equals(TaskDefinition that) {
    if (that == null)
      return false;

    boolean this_present_command = true && this.isSetCommand();
    boolean that_present_command = true && that.isSetCommand();
    if (this_present_command || that_present_command) {
      if (!(this_present_command && that_present_command))
        return false;
      if (!this.command.equals(that.command))
        return false;
    }

    boolean this_present_snapshotId = true && this.isSetSnapshotId();
    boolean that_present_snapshotId = true && that.isSetSnapshotId();
    if (this_present_snapshotId || that_present_snapshotId) {
      if (!(this_present_snapshotId && that_present_snapshotId))
        return false;
      if (!this.snapshotId.equals(that.snapshotId))
        return false;
    }

    boolean this_present_taskId = true && this.isSetTaskId();
    boolean that_present_taskId = true && that.isSetTaskId();
    if (this_present_taskId || that_present_taskId) {
      if (!(this_present_taskId && that_present_taskId))
        return false;
      if (!this.taskId.equals(that.taskId))
        return false;
    }

    boolean this_present_timeoutMs = true && this.isSetTimeoutMs();
    boolean that_present_timeoutMs = true && that.isSetTimeoutMs();
    if (this_present_timeoutMs || that_present_timeoutMs) {
      if (!(this_present_timeoutMs && that_present_timeoutMs))
        return false;
      if (this.timeoutMs != that.timeoutMs)
        return false;
    }

    return true;
  }

  @Override
  public int hashCode() {
    List<Object> list = new ArrayList<Object>();

    boolean present_command = true && (isSetCommand());
    list.add(present_command);
    if (present_command)
      list.add(command);

    boolean present_snapshotId = true && (isSetSnapshotId());
    list.add(present_snapshotId);
    if (present_snapshotId)
      list.add(snapshotId);

    boolean present_taskId = true && (isSetTaskId());
    list.add(present_taskId);
    if (present_taskId)
      list.add(taskId);

    boolean present_timeoutMs = true && (isSetTimeoutMs());
    list.add(present_timeoutMs);
    if (present_timeoutMs)
      list.add(timeoutMs);

    return list.hashCode();
  }

  @Override
  public int compareTo(TaskDefinition other) {
    if (!getClass().equals(other.getClass())) {
      return getClass().getName().compareTo(other.getClass().getName());
    }

    int lastComparison = 0;

    lastComparison = Boolean.valueOf(isSetCommand()).compareTo(other.isSetCommand());
    if (lastComparison != 0) {
      return lastComparison;
    }
    if (isSetCommand()) {
      lastComparison = org.apache.thrift.TBaseHelper.compareTo(this.command, other.command);
      if (lastComparison != 0) {
        return lastComparison;
      }
    }
    lastComparison = Boolean.valueOf(isSetSnapshotId()).compareTo(other.isSetSnapshotId());
    if (lastComparison != 0) {
      return lastComparison;
    }
    if (isSetSnapshotId()) {
      lastComparison = org.apache.thrift.TBaseHelper.compareTo(this.snapshotId, other.snapshotId);
      if (lastComparison != 0) {
        return lastComparison;
      }
    }
    lastComparison = Boolean.valueOf(isSetTaskId()).compareTo(other.isSetTaskId());
    if (lastComparison != 0) {
      return lastComparison;
    }
    if (isSetTaskId()) {
      lastComparison = org.apache.thrift.TBaseHelper.compareTo(this.taskId, other.taskId);
      if (lastComparison != 0) {
        return lastComparison;
      }
    }
    lastComparison = Boolean.valueOf(isSetTimeoutMs()).compareTo(other.isSetTimeoutMs());
    if (lastComparison != 0) {
      return lastComparison;
    }
    if (isSetTimeoutMs()) {
      lastComparison = org.apache.thrift.TBaseHelper.compareTo(this.timeoutMs, other.timeoutMs);
      if (lastComparison != 0) {
        return lastComparison;
      }
    }
    return 0;
  }

  public _Fields fieldForId(int fieldId) {
    return _Fields.findByThriftId(fieldId);
  }

  public void read(org.apache.thrift.protocol.TProtocol iprot) throws org.apache.thrift.TException {
    schemes.get(iprot.getScheme()).getScheme().read(iprot, this);
  }

  public void write(org.apache.thrift.protocol.TProtocol oprot) throws org.apache.thrift.TException {
    schemes.get(oprot.getScheme()).getScheme().write(oprot, this);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder("TaskDefinition(");
    boolean first = true;

    sb.append("command:");
    if (this.command == null) {
      sb.append("null");
    } else {
      sb.append(this.command);
    }
    first = false;
    if (isSetSnapshotId()) {
      if (!first) sb.append(", ");
      sb.append("snapshotId:");
      if (this.snapshotId == null) {
        sb.append("null");
      } else {
        sb.append(this.snapshotId);
      }
      first = false;
    }
    if (isSetTaskId()) {
      if (!first) sb.append(", ");
      sb.append("taskId:");
      if (this.taskId == null) {
        sb.append("null");
      } else {
        sb.append(this.taskId);
      }
      first = false;
    }
    if (isSetTimeoutMs()) {
      if (!first) sb.append(", ");
      sb.append("timeoutMs:");
      sb.append(this.timeoutMs);
      first = false;
    }
    sb.append(")");
    return sb.toString();
  }

  public void validate() throws org.apache.thrift.TException {
    // check for required fields
    if (command == null) {
      throw new org.apache.thrift.protocol.TProtocolException("Required field 'command' was not present! Struct: " + toString());
    }
    // check for sub-struct validity
    if (command != null) {
      command.validate();
    }
  }

  private void writeObject(java.io.ObjectOutputStream out) throws java.io.IOException {
    try {
      write(new org.apache.thrift.protocol.TCompactProtocol(new org.apache.thrift.transport.TIOStreamTransport(out)));
    } catch (org.apache.thrift.TException te) {
      throw new java.io.IOException(te);
    }
  }

  private void readObject(java.io.ObjectInputStream in) throws java.io.IOException, ClassNotFoundException {
    try {
      // it doesn't seem like you should have to do this, but java serialization is wacky, and doesn't call the default constructor.
      __isset_bitfield = 0;
      read(new org.apache.thrift.protocol.TCompactProtocol(new org.apache.thrift.transport.TIOStreamTransport(in)));
    } catch (org.apache.thrift.TException te) {
      throw new java.io.IOException(te);
    }
  }

  private static class TaskDefinitionStandardSchemeFactory implements SchemeFactory {
    public TaskDefinitionStandardScheme getScheme() {
      return new TaskDefinitionStandardScheme();
    }
  }

  private static class TaskDefinitionStandardScheme extends StandardScheme<TaskDefinition> {

    public void read(org.apache.thrift.protocol.TProtocol iprot, TaskDefinition struct) throws org.apache.thrift.TException {
      org.apache.thrift.protocol.TField schemeField;
      iprot.readStructBegin();
      while (true)
      {
        schemeField = iprot.readFieldBegin();
        if (schemeField.type == org.apache.thrift.protocol.TType.STOP) { 
          break;
        }
        switch (schemeField.id) {
          case 1: // COMMAND
            if (schemeField.type == org.apache.thrift.protocol.TType.STRUCT) {
              struct.command = new Command();
              struct.command.read(iprot);
              struct.setCommandIsSet(true);
            } else { 
              org.apache.thrift.protocol.TProtocolUtil.skip(iprot, schemeField.type);
            }
            break;
          case 2: // SNAPSHOT_ID
            if (schemeField.type == org.apache.thrift.protocol.TType.STRING) {
              struct.snapshotId = iprot.readString();
              struct.setSnapshotIdIsSet(true);
            } else { 
              org.apache.thrift.protocol.TProtocolUtil.skip(iprot, schemeField.type);
            }
            break;
          case 3: // TASK_ID
            if (schemeField.type == org.apache.thrift.protocol.TType.STRING) {
              struct.taskId = iprot.readString();
              struct.setTaskIdIsSet(true);
            } else { 
              org.apache.thrift.protocol.TProtocolUtil.skip(iprot, schemeField.type);
            }
            break;
          case 4: // TIMEOUT_MS
            if (schemeField.type == org.apache.thrift.protocol.TType.I32) {
              struct.timeoutMs = iprot.readI32();
              struct.setTimeoutMsIsSet(true);
            } else { 
              org.apache.thrift.protocol.TProtocolUtil.skip(iprot, schemeField.type);
            }
            break;
          default:
            org.apache.thrift.protocol.TProtocolUtil.skip(iprot, schemeField.type);
        }
        iprot.readFieldEnd();
      }
      iprot.readStructEnd();

      // check for required fields of primitive type, which can't be checked in the validate method
      struct.validate();
    }

    public void write(org.apache.thrift.protocol.TProtocol oprot, TaskDefinition struct) throws org.apache.thrift.TException {
      struct.validate();

      oprot.writeStructBegin(STRUCT_DESC);
      if (struct.command != null) {
        oprot.writeFieldBegin(COMMAND_FIELD_DESC);
        struct.command.write(oprot);
        oprot.writeFieldEnd();
      }
      if (struct.snapshotId != null) {
        if (struct.isSetSnapshotId()) {
          oprot.writeFieldBegin(SNAPSHOT_ID_FIELD_DESC);
          oprot.writeString(struct.snapshotId);
          oprot.writeFieldEnd();
        }
      }
      if (struct.taskId != null) {
        if (struct.isSetTaskId()) {
          oprot.writeFieldBegin(TASK_ID_FIELD_DESC);
          oprot.writeString(struct.taskId);
          oprot.writeFieldEnd();
        }
      }
      if (struct.isSetTimeoutMs()) {
        oprot.writeFieldBegin(TIMEOUT_MS_FIELD_DESC);
        oprot.writeI32(struct.timeoutMs);
        oprot.writeFieldEnd();
      }
      oprot.writeFieldStop();
      oprot.writeStructEnd();
    }

  }

  private static class TaskDefinitionTupleSchemeFactory implements SchemeFactory {
    public TaskDefinitionTupleScheme getScheme() {
      return new TaskDefinitionTupleScheme();
    }
  }

  private static class TaskDefinitionTupleScheme extends TupleScheme<TaskDefinition> {

    @Override
    public void write(org.apache.thrift.protocol.TProtocol prot, TaskDefinition struct) throws org.apache.thrift.TException {
      TTupleProtocol oprot = (TTupleProtocol) prot;
      struct.command.write(oprot);
      BitSet optionals = new BitSet();
      if (struct.isSetSnapshotId()) {
        optionals.set(0);
      }
      if (struct.isSetTaskId()) {
        optionals.set(1);
      }
      if (struct.isSetTimeoutMs()) {
        optionals.set(2);
      }
      oprot.writeBitSet(optionals, 3);
      if (struct.isSetSnapshotId()) {
        oprot.writeString(struct.snapshotId);
      }
      if (struct.isSetTaskId()) {
        oprot.writeString(struct.taskId);
      }
      if (struct.isSetTimeoutMs()) {
        oprot.writeI32(struct.timeoutMs);
      }
    }

    @Override
    public void read(org.apache.thrift.protocol.TProtocol prot, TaskDefinition struct) throws org.apache.thrift.TException {
      TTupleProtocol iprot = (TTupleProtocol) prot;
      struct.command = new Command();
      struct.command.read(iprot);
      struct.setCommandIsSet(true);
      BitSet incoming = iprot.readBitSet(3);
      if (incoming.get(0)) {
        struct.snapshotId = iprot.readString();
        struct.setSnapshotIdIsSet(true);
      }
      if (incoming.get(1)) {
        struct.taskId = iprot.readString();
        struct.setTaskIdIsSet(true);
      }
      if (incoming.get(2)) {
        struct.timeoutMs = iprot.readI32();
        struct.setTimeoutMsIsSet(true);
      }
    }
  }

}
