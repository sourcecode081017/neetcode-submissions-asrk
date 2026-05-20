class Solution {
    public int numUniqueEmails(String[] emails) {
        Set<String> set = new HashSet<>();
        for(String email : emails) {
            set.add(processEmailAddr(email));
        }
        return set.size();
    }
    private String processEmailAddr(String email) {
        String emailSplit[] = email.split("@");
        String name = emailSplit[0];
        String domain = emailSplit[1];
        StringBuilder nameBuilder = new StringBuilder("");
        for( char c : name.toCharArray()) {
            if(c != '.') {
                nameBuilder.append(c);
            }
            if(c == '+') {
                break;
            }
        }
        return nameBuilder.toString() + domain;
    }
}